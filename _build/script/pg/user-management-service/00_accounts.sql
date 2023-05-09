\c usermanagementdb

BEGIN;
DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts
(
   account_id SERIAL NOT NULL,
   first_name text,
   last_name text,
   email text,
   address text,
   phone_number text,
   role int,
   hash_text text,
   open_date TIMESTAMP (0) WITH TIME ZONE,
   PRIMARY KEY (account_id)
);
END;