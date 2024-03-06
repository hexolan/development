CREATE TABLE items (
    id bigserial PRIMARY KEY,
    name varchar(256) NOT NULL,

    created_at timestamp NOT NULL DEFAULT timezone('utc', now()),
    updated_at timestamp
);

CREATE TABLE event_outbox (
    id bigserial PRIMARY KEY,

    aggregatetype varchar(128) NOT NULL,
    aggregateid varchar(128) NOT NULL,
    payload bytea NOT NULL
);