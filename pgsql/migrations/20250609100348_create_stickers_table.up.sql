BEGIN;

CREATE TABLE stickers (
    id varchar(64) NOT NULL PRIMARY KEY,
    active boolean NOT NULL DEFAULT false,
    name varchar(255) NULL,
    description varchar(255) NULL,
    image_url varchar(255) NULL,
    show_phone_number boolean NOT NULL DEFAULT false,
    phone_number varchar(255) NULL,
    show_email boolean NOT NULL DEFAULT false,
    email varchar(255) NULL,
    show_instagram boolean NOT NULL DEFAULT false,
    instagram_url varchar(255) NULL,
    show_facebook boolean NOT NULL DEFAULT false,
    facebook_url varchar(255) NULL,
    user_id varchar(64) NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);



COMMIT;