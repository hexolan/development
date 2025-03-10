CREATE TABLE items (
    id bigserial PRIMARY KEY,
    name varchar(256) NOT NULL,

    created_at timestamp NOT NULL DEFAULT timezone('utc', now()),
    updated_at timestamp
);

CREATE TABLE event_outbox (
    id bigserial PRIMARY KEY,

    topic varchar(128) NOT NULL,
    msg bytea NOT NULL
);