BEGIN;

CREATE TABLE messages (
    id varchar(64) NOT NULL PRIMARY KEY,
    sticker_id varchar(64) NOT NULL,
    message text NOT NULL,
    urgency_level varchar(10) NOT NULL DEFAULT 'normal',
    created_at timestamp NOT NULL
);

COMMIT;