
-- Run marketplace2.sql in marketplace2 container first , then run this file
-- ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

CREATE EXTENSION postgres_fdw;
CREATE SERVER node1 FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host 'marketplace2', dbname 'ds_db');
CREATE USER MAPPING FOR postgres SERVER node1 OPTIONS (user 'postgres' , password '1700455');


-- User table

CREATE TABLE "users" ("id" bigserial,"email" text,"password" text,"name" text,"balance" decimal,"image_url" text,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
CREATE INDEX "idx_users_id" ON "users" ("id");

CREATE TABLE Users_0 (
CHECK ( id%2=0)
) INHERITS (Users);

CREATE FOREIGN TABLE Users_1 (
CHECK ( id%2=1)
) INHERITS (Users) server node1;

CREATE OR REPLACE FUNCTION users_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
IF ( NEW.id % 2 = 0) THEN
INSERT INTO Users_0 VALUES (NEW.*);
ELSIF ( NEW.id % 2 = 1 ) THEN
INSERT INTO Users_1 VALUES (NEW.*);
END IF;
RETURN NULL;
END;

$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_users_trigger
    BEFORE INSERT ON Users
    FOR EACH ROW EXECUTE FUNCTION users_insert_trigger();


-- Product table

CREATE TABLE "products" ("id" bigserial,"user_id" bigint,"title" text,"content" text,"image_url" text,"price" decimal,"status" boolean,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_products_deleted_at" ON "products" ("deleted_at");
CREATE INDEX "idx_products_id" ON "products" ("id");

CREATE TABLE Products_0 (
CHECK ( id%2=0)
) INHERITS (Products);

CREATE FOREIGN TABLE Products_1 (
CHECK ( id%2=1)
) INHERITS (Products) server node1;


CREATE OR REPLACE FUNCTION products_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
IF ( NEW.id % 2 = 0) THEN
INSERT INTO Products_0 VALUES (NEW.*);
ELSIF ( NEW.id % 2 = 1 ) THEN
INSERT INTO Products_1 VALUES (NEW.*);
END IF;
RETURN NULL;
END;

$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_products_trigger
    BEFORE INSERT ON Products
    FOR EACH ROW EXECUTE FUNCTION products_insert_trigger();




-- Stores table

CREATE TABLE "stores" ("id" bigserial,"title" text,"user_id" bigint,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_stores_deleted_at" ON "stores" ("deleted_at");
CREATE INDEX "idx_stores_id" ON "stores" ("deleted_at");

CREATE TABLE Stores_0 (
CHECK ( id%2=0)
) INHERITS (Stores);

CREATE FOREIGN TABLE Stores_1 (
CHECK ( id%2=1)
) INHERITS (Stores) server node1;



CREATE OR REPLACE FUNCTION stores_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
IF ( NEW.id % 2 = 0) THEN
INSERT INTO Stores_0 VALUES (NEW.*);
ELSIF ( NEW.id % 2 = 1 ) THEN
INSERT INTO Stores_1 VALUES (NEW.*);
END IF;
RETURN NULL;
END;

$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_stores_trigger
    BEFORE INSERT ON Stores
    FOR EACH ROW EXECUTE FUNCTION stores_insert_trigger();



CREATE TABLE "product_store" ("store_id" bigint,"product_id" bigint);
CREATE INDEX "idx_product_store_store_id" ON "product_store" ("store_id");
CREATE INDEX "idx_product_store_product_id" ON "product_store" ("product_id");





-- Orders table

CREATE TABLE "orders" ("id" bigserial,"buyer_id" bigint,"seller_id" bigint,"product_id" bigint,"price" decimal,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_orders_deleted_at" ON "orders" ("deleted_at");
CREATE INDEX "idx_orders_id" ON "orders" ("id");

CREATE TABLE Orders_0 (
CHECK ( id%2=0)
) INHERITS (Orders);

CREATE FOREIGN TABLE Orders_1 (
CHECK ( id%2=1)
) INHERITS (Orders) server node1;


CREATE OR REPLACE FUNCTION orders_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
IF ( NEW.id % 2 = 0) THEN
INSERT INTO Orders_0 VALUES (NEW.*);
ELSIF ( NEW.id % 2 = 1 ) THEN
INSERT INTO Orders_1 VALUES (NEW.*);
END IF;
RETURN NULL;
END;

$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_orders_trigger
    BEFORE INSERT ON Orders
    FOR EACH ROW EXECUTE FUNCTION orders_insert_trigger();




-- Deposits table

CREATE TABLE "deposits" ("id" bigserial,"user_id" bigint,"balance_before" decimal,"amount" decimal,"balance_after" decimal,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_deposits_deleted_at" ON "deposits" ("deleted_at");
CREATE INDEX "idx_deposits_id" ON "deposits" ("id");

CREATE TABLE Deposits_0 (
CHECK ( id%2=0)
) INHERITS (Deposits);

CREATE FOREIGN TABLE Deposits_1 (
CHECK ( id%2=1)
) INHERITS (Deposits) server node1;


CREATE OR REPLACE FUNCTION deposits_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
IF ( NEW.id % 2 = 0) THEN
INSERT INTO Deposits_0 VALUES (NEW.*);
ELSIF ( NEW.id % 2 = 1 ) THEN
INSERT INTO Deposits_1 VALUES (NEW.*);
END IF;
RETURN NULL;
END;

$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_deposits_trigger
    BEFORE INSERT ON Deposits
    FOR EACH ROW EXECUTE FUNCTION deposits_insert_trigger();


