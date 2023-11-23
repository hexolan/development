CREATE TABLE shipping_orders (
    order_id varchar(64) PRIMARY KEY,
    status smallint NOT NULL,
);
CREATE TABLE shipping_order_items (
    order_id varchar(64),
    product_id varchar(64),

    quantity integer NOT NULL,

    PRIMARY KEY (order_id, product_id)
);