\c authenticationdb

BEGIN;
DROP TABLE IF EXISTS product_details;
CREATE TABLE product_details
(
   product_detail_id BIGINT NOT NULL,
   product_id BIGINT NOT NULL,
   product_description text,
   PRIMARY KEY (product_detail_id)
);
END;