BEGIN;

CREATE TABLE items (
    id bigserial PRIMARY KEY,
    status smallint NOT NULL,

    title varchar(128) NOT NULL,

    created_at timestamp NOT NULL DEFAULT timezone('utc', now())
);

END;