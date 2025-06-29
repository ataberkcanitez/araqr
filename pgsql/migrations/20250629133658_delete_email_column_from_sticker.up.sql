BEGIN;

alter table stickers
    drop column show_email;

alter table stickers
    drop column email;

COMMIT;