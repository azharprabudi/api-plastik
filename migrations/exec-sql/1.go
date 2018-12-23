package execsql

// First ...
var First = `
CREATE TABLE "item_categories" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT item_categories_pk PRIMARY KEY ("id")
);

CREATE TABLE "items" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"category_id" uuid NOT NULL,
	"unit_id" uuid NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY ("id")
);

CREATE TABLE "item_units" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT item_units_pk PRIMARY KEY ("id")
);

INSERT INTO "item_units"(id, name, created_at)
		VALUES ('22641d60-c7e9-413d-b0d7-59618810a3c4', 'kg', CURRENT_TIMESTAMP)
		ON CONFLICT DO NOTHING;

INSERT INTO "item_units"(id, name, created_at)
	VALUES ('1c4d74e7-cd09-41f0-8190-5a2dcaffb195', 'pcs', CURRENT_TIMESTAMP)
	ON CONFLICT DO NOTHING;

ALTER TABLE items
   ADD CONSTRAINT fk_item_categories
   FOREIGN KEY ("category_id") 
   REFERENCES "item_categories"("id");

ALTER TABLE items
   ADD CONSTRAINT fk_item_units
   FOREIGN KEY ("unit_id") 
   REFERENCES "item_units"("id");

CREATE TABLE "suppliers" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"phone" varchar(15) NULL,
	"address" text NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT suppliers_pk PRIMARY KEY ("id")
);

CREATE TABLE "sellers" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"phone" varchar(15) NULL,
	"address" text NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT sellers_pk PRIMARY KEY ("id")
);

CREATE TABLE "transactions" (
	"id" uuid NOT NULL,
	"note" text,
	"user_id" uuid NOT NULL,
	"company_id" uuid NOT NULL,
	"type" varchar(20) NOT NULL,
	"amount" numeric(20, 2) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT transactions_pk PRIMARY KEY ("id")
);

CREATE TABLE "transactions_in" (
	"id" uuid NOT NULL,
	"transaction_id" uuid NOT NULL,
	"supplier_id" uuid NOT NULL,
	CONSTRAINT transactions_in_pk PRIMARY KEY ("id")
);

CREATE TABLE "transactions_out" (
	"id" uuid NOT NULL,
	"transaction_id" uuid NOT NULL,
	"seller_id" uuid NOT NULL,
	CONSTRAINT transactions_out_pk PRIMARY KEY ("id")
);

CREATE TABLE "transactions_etc" (
	"id" uuid NOT NULL,
	"transaction_id" uuid NOT NULL,
	"transaction_etc_type" uuid NOT NULL,
	CONSTRAINT transactions_etc_pk PRIMARY KEY ("id")
);

CREATE TABLE "transaction_etc_types" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT transaction_etc_types_pk PRIMARY KEY ("id")
);

ALTER TABLE transactions_in
   ADD CONSTRAINT fk_transactions_in
   FOREIGN KEY ("transaction_id") 
   REFERENCES "transactions"("id");

ALTER TABLE transactions_in
	ADD CONSTRAINT fk_suppliers
	FOREIGN KEY ("supplier_id") 
	REFERENCES "suppliers"("id");

ALTER TABLE transactions_out
	ADD CONSTRAINT fk_transactions_out
	FOREIGN KEY ("transaction_id") 
	REFERENCES "transactions"("id");
 
ALTER TABLE transactions_out
	ADD CONSTRAINT fk_sellers
	FOREIGN KEY ("seller_id") 
	REFERENCES "sellers"("id");

ALTER TABLE transactions_etc
	ADD CONSTRAINT fk_transactions
	FOREIGN KEY ("transaction_id") 
	REFERENCES "transactions"("id");

ALTER TABLE transactions_etc
	ADD CONSTRAINT fk_transaction_etc_types
	FOREIGN KEY ("transaction_etc_type") 
	REFERENCES "transaction_etc_types"("id");
 
CREATE TABLE "transaction_details" (
	"id" uuid NOT NULL,
	"item_id" uuid NOT NULL,
	"item_name" varchar(100) NOT NULL,
	"transaction_id" uuid NOT NULL,
	"qty" int NOT NULL,
	"amount" numeric(20, 2) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT transaction_details_pk PRIMARY KEY ("id")
);

ALTER TABLE transaction_details
	ADD CONSTRAINT fk_item
	FOREIGN KEY ("item_id") 
	REFERENCES "items"("id");

ALTER TABLE transaction_details
	ADD CONSTRAINT fk_transaction
	FOREIGN KEY ("transaction_id") 
	REFERENCES "transactions"("id");

CREATE TABLE "transaction_images" (
	"id" uuid NOT NULL,
	"image" varchar(100) NOT NULL,
	"transaction_id" uuid NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT transaction_images_pk PRIMARY KEY ("id")
);

ALTER TABLE transaction_images
	ADD CONSTRAINT fk_transaction
	FOREIGN KEY ("transaction_id") 
	REFERENCES "transactions"("id");

CREATE TABLE "item_stock_logs" (
	"id" uuid NOT NULL,
	"item_name" varchar(100),
	"qty" int NOT NULL,
	"item_id" uuid NOT NULL,
	"transaction_id" uuid NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT item_stock_logs_pk PRIMARY KEY ("id")
);

ALTER TABLE item_stock_logs
	ADD CONSTRAINT fk_transaction
	FOREIGN KEY ("transaction_id") 
	REFERENCES "transactions"("id");

`
