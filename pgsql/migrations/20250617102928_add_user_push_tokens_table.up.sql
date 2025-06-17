CREATE TABLE user_push_tokens
(
    id          VARCHAR(64) PRIMARY KEY,
    user_id     VARCHAR(64),
    push_token  VARCHAR(255) UNIQUE NOT NULL,
    platform    VARCHAR(10)         NOT NULL, -- 'ios' or 'android'
    device_name VARCHAR(255),
    os_name     VARCHAR(50),
    os_version  VARCHAR(50),
    is_active   BOOLEAN   DEFAULT true,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);