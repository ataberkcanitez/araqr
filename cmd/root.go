package cmd

import (
	"github.com/ataberkcanitez/araqr/cmd/serve"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RunRootCmd entrypoint of the root command
func RunRootCmd() error {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cmd := &cobra.Command{
		Use:   "araqr",
		Short: "araqr application",
		Long:  "araqr application",
	}

	cmd.PersistentFlags().String("app.name", "araqr", "Application name")
	cmd.PersistentFlags().String("app.env", "dev", "Application environment")

	cmd.PersistentFlags().String("log.encoding", "json", "Log format (json, console)")
	cmd.PersistentFlags().String("log.level", "info", "Log level")

	cmd.AddCommand(serve.New())

	_ = viper.BindPFlags(cmd.PersistentFlags())

	return cmd.Execute()
}
