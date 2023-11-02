CREATE TABLE transactions (
    id bigserial PRIMARY KEY,

    order_id varchar(64) NOT NULL,
    customer_id varchar(64),

    amount money NOT NULL,

    executed_at timestamp,
    reversed_at timestamp
);