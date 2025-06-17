BEGIN;

ALTER TABLE messages
    DROP COLUMN read;

ALTER TABLE messages
    DROP COLUMN updated_at;

COMMIT;