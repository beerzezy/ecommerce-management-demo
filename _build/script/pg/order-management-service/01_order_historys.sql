\c authenticationdb

BEGIN;
DROP TABLE IF EXISTS order_historys;
CREATE TABLE order_historys
(
   order_historys_id BIGINT NOT NULL,
   account_id BIGINT NOT NULL,
   product_id BIGINT NOT NULL,
   order_date TIMESTAMP (0) WITH TIME ZONE,
   PRIMARY KEY (orders_id)
);
END;
