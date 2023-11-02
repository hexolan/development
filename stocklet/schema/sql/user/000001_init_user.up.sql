CREATE TABLE users (
    id bigserial PRIMARY KEY,

    first_name varchar(64) NOT NULL,
    last_name varchar(64) NOT NULL,

    /* todo: address lines */

    created_at timestamp NOT NULL DEFAULT timezone('utc', now()),
    updated_at timestamp
);