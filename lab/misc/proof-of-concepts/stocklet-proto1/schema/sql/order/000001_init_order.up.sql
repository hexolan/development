CREATE TABLE orders (
    id bigserial PRIMARY KEY,
    status smallint NOT NULL,

    customer_id varchar(64) NOT NULL,
    transaction_id varchar(64),

    created_at timestamp NOT NULL DEFAULT timezone('utc', now()),
    updated_at timestamp
);

CREATE TABLE order_items (
    order_id bigserial,
    product_id varchar(64),

    quantity integer NOT NULL,

    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);

CREATE TABLE event_outbox (
    id bigserial PRIMARY KEY,

    topic varchar(128) NOT NULL,
    msg bytea NOT NULL
);