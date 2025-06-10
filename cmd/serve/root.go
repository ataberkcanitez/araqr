package serve

import (
	"context"
	messagesRepo "github.com/ataberkcanitez/araqr/internal/adapter/repository/message"
	"github.com/ataberkcanitez/araqr/internal/adapter/repository/refresh_token"
	stickerRepo "github.com/ataberkcanitez/araqr/internal/adapter/repository/sticker"
	"github.com/ataberkcanitez/araqr/internal/adapter/repository/user"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/services/auth"
	"github.com/ataberkcanitez/araqr/internal/application/services/sticker"
	"github.com/ataberkcanitez/araqr/log"
	"github.com/ataberkcanitez/araqr/pgsql"
	pgsql2 "github.com/ataberkcanitez/araqr/pgsql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// New creates a new serve command
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve REST API",
		Long:  "Serve REST API",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return viper.BindPFlags(cmd.PersistentFlags())
		},
		RunE: runServeHTTP,
		Args: cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().String("server.allow-origin", "*", "Allow origin")
	cmd.PersistentFlags().String("server.port", ":8080", "HTTP port")

	cmd.PersistentFlags().String("auth.secret-key", "", "secret key")
	cmd.PersistentFlags().Duration("auth.access-token-expiry", 15*time.Minute, "access token expiry")
	cmd.PersistentFlags().Duration("auth.refresh-token-expiry", 7*24*time.Hour, "refresh token expiry")

	cmd.PersistentFlags().StringP("db.host", "H", "localhost", "database host")
	cmd.PersistentFlags().IntP("db.port", "P", 5432, "database port")
	cmd.PersistentFlags().StringP("db.dbname", "d", "araqr", "database name")
	cmd.PersistentFlags().StringP("db.user", "U", "postgres", "database user")
	cmd.PersistentFlags().String("db.password", "toor", "database password")
	cmd.PersistentFlags().String("db.sslmode", "require", "database sslmode")

	cmd.PersistentFlags().String("email.sender-address", "noreply@dev.araqr.com", "sender email address")
	cmd.PersistentFlags().String("email.support.email-address", "destek@dev.randevumu.com", "support email address")
	cmd.PersistentFlags().String("email.support.phone-number", "+90 555 555 55 55", "support phone number")
	cmd.PersistentFlags().String("email.support.website-link", "https://dev.araqr.com", "support website link")

	cmd.PersistentFlags().String("storage.bucket", "araqr-<TODO:update>", "aws s3 bucket")

	return cmd
}

type serveConfig struct {
	Server struct {
		AllowOrigin string `mapstructure:"allow-origin"`
		Port        string `mapstructure:"port"`
	} `mapstructure:"server"`
	Log struct {
		Encoding string `mapstructure:"encoding"`
		Level    string `mapstructure:"level"`
	} `mapstructure:"log"`
	DB   pgsql2.Config `mapstructure:"db"`
	Auth auth.Config   `mapstructure:"auth"`
}

func runServeHTTP(cmd *cobra.Command, _ []string) error {
	var cfg serveConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}

	log.Init(cfg.Log.Encoding, cfg.Log.Level)

	ctx := cmd.Context()

	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, os.Interrupt)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.Server.AllowOrigin},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodHead},
	}))
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Skipper:   middleware.DefaultRequestIDConfig.Skipper,
		Generator: middleware.DefaultRequestIDConfig.Generator,
		RequestIDHandler: func(c echo.Context, requestID string) {
			ctx := context.WithValue(c.Request().Context(),
				"request_id", requestID, // nolint:staticcheck
			)
			c.SetRequest(c.Request().WithContext(ctx))
		},
		TargetHeader: middleware.DefaultRequestIDConfig.TargetHeader,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/health"
		},
		Format: `{"time":"${time_rfc3339_nano}","request_id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: middleware.DefaultLoggerConfig.CustomTimeFormat,
		CustomTagFunc:    middleware.DefaultLoggerConfig.CustomTagFunc,
		Output:           middleware.DefaultLoggerConfig.Output,
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.HTTPErrorHandler = web.ErrorHandler

	handlers, err := createHandlers(ctx, &cfg)
	if err != nil {
		return err
	}

	for _, h := range handlers {
		h.RegisterRoutes(e)
	}

	go func() {
		if err := e.Start(cfg.Server.Port); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error(ctx, "Error occurred while listening.", err)
		}
	}()

	<-quit

	log.Info(ctx, "Shutting down the server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error(ctx, "Error occurred while shutting down.", err)
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Server shutdown timed out.")
	default:
		log.Info(ctx, "Server shutdown gracefully.")
	}

	return nil
}

type serverHandler interface {
	RegisterRoutes(e *echo.Echo)
}

func createHandlers(ctx context.Context, cfg *serveConfig) ([]serverHandler, error) {
	if err := pgsql.MigrateUp(&cfg.DB); err != nil {
		return nil, err
	}

	db, err := pgsql2.Connect(ctx, &cfg.DB)
	if err != nil {
		return nil, err
	}

	userPg := user.NewRepository(db)
	stickerRepository := stickerRepo.NewRepository(db)
	refreshTokenPg := refresh_token.NewRepository(db)
	messageRepository := messagesRepo.NewMessageRepository(db)

	authSvc := auth.NewService(&cfg.Auth, userPg, refreshTokenPg)
	stickerSvc := sticker.NewService(userPg, stickerRepository, messageRepository)

	return []serverHandler{
		web.NewAuthHandler(authSvc),
		web.NewStickerHandler(stickerSvc, authSvc),
	}, nil
}
