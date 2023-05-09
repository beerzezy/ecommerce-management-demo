\c authenticationdb

BEGIN;
DROP TABLE IF EXISTS orders;
CREATE TABLE orders
(
   orders_id BIGINT NOT NULL,
   account_id BIGINT NOT NULL,
   product_id BIGINT NOT NULL,
   product_name text,
   product_description text,
   order_date TIMESTAMP (0) WITH TIME ZONE,
   PRIMARY KEY (orders_id)
);
END;