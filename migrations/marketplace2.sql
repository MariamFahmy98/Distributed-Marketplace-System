CREATE TABLE "users" ("id" bigserial,"email" text,"password" text,"name" text,"balance" decimal,"image_url" text,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
CREATE INDEX "idx_users_id" ON "users" ("id");

CREATE  TABLE Users_1 (
CHECK ( id%2=1)
) INHERITS (Users);





CREATE TABLE "products" ("id" bigserial,"user_id" bigint,"title" text,"content" text,"image_url" text,"price" decimal,"status" boolean,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_products_deleted_at" ON "products" ("deleted_at");
CREATE INDEX "idx_products_id" ON "products" ("id");

CREATE  TABLE Products_1 (
CHECK ( id%2=1)
) INHERITS (Products);





CREATE TABLE "stores" ("id" bigserial,"title" text,"user_id" bigint,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_stores_deleted_at" ON "stores" ("deleted_at");
CREATE INDEX "idx_stores_id" ON "stores" ("deleted_at");

CREATE  TABLE Stores_1 (
CHECK ( id%2=1)
) INHERITS (Stores);





CREATE TABLE "product_store" ("store_id" bigint,"product_id" bigint);
CREATE INDEX "idx_product_store_store_id" ON "product_store" ("store_id");
CREATE INDEX "idx_product_store_product_id" ON "product_store" ("product_id");





CREATE TABLE "orders" ("id" bigserial,"buyer_id" bigint,"seller_id" bigint,"product_id" bigint,"price" decimal,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_orders_deleted_at" ON "orders" ("deleted_at");
CREATE INDEX "idx_orders_id" ON "orders" ("id");

CREATE  TABLE Orders_1 (
CHECK ( id%2=1)
) INHERITS (Orders);





CREATE TABLE "deposits" ("id" bigserial,"user_id" bigint,"balance_before" decimal,"amount" decimal,"balance_after" decimal,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz);
CREATE INDEX "idx_deposits_deleted_at" ON "deposits" ("deleted_at");
CREATE INDEX "idx_deposits_id" ON "deposits" ("id");

CREATE  TABLE Deposits_1 (
CHECK ( id%2=1)
) INHERITS (Deposits);

