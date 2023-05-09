\c authenticationdb

BEGIN;
DROP TABLE IF EXISTS products;
CREATE TABLE products
(
   product_id BIGINT NOT NULL,
   product_name text,
   product_description text,
   image_ref text,
   remaining text,
   in_stock boolean,
   PRIMARY KEY (product_id)
);
END;