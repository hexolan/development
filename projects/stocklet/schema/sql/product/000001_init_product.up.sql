CREATE TABLE products (
    id bigserial PRIMARY KEY,

    name varchar(64) NOT NULL,
    description TEXT,

    unit_price money NOT NULL,

    created_at timestamp NOT NULL DEFAULT timezone('utc', now()),
    updated_at timestamp
);