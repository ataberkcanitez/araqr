BEGIN;

CREATE TABLE users (
    id VARCHAR(64) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_users_email ON users (email);

CREATE TABLE refresh_tokens (
    user_id VARCHAR(64) NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL,
    valid BOOLEAN NOT NULL DEFAULT TRUE,
    valid_until TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);

CREATE INDEX idx_refresh_tokens_token ON refresh_tokens (token);

CREATE TABLE password_reset_tokens (
    user_id VARCHAR(64) NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token VARCHAR(64) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_password_reset_tokens_token ON password_reset_tokens (token);
CREATE INDEX idx_password_reset_tokens_user_id ON password_reset_tokens (user_id);

END;
