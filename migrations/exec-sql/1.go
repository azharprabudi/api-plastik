package execsql

import "github.com/azharprabudi/api-plastik/internal/user/value"

// First ...
var First = `
CREATE TABLE "companies" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"type" varchar(50) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT companies_pk PRIMARY KEY ("id")
);

INSERT INTO "companies"(id, name, type, created_at) VALUES ('` + value.COMPANY_ID.String() + `', 'PT. Berkah Jaya Plastik', '` + value.RETAIL + `', CURRENT_TIMESTAMP) ON CONFLICT DO NOTHING;

CREATE TABLE "user_groups" (
	"id" uuid NOT NULL,
	"name" varchar(100) NOT NULL, 
	"created_at" timestamptz NOT NULL,
	CONSTRAINT user_groups_pk PRIMARY KEY ("id")
);

INSERT INTO "user_groups"(id, name, created_at)
		VALUES ('` + value.GROUP_ID.String() + `', 'admin', CURRENT_TIMESTAMP)
		ON CONFLICT DO NOTHING;

CREATE TABLE "users" (
	"id" uuid NOT NULL,
	"username" varchar(100) NOT NULL, 
	"name" varchar(100) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY ("id")
);

ALTER TABLE users 
ADD CONSTRAINT users_unique_name 
UNIQUE (username);

INSERT INTO "users"(id, username, name, created_at)
		VALUES ('` + value.USER_ID.String() + `', 'admin_plastik', 'admin', CURRENT_TIMESTAMP)
		ON CONFLICT DO NOTHING;

CREATE TABLE "users_company" (
	"user_id" uuid NOT NULL,
	"company_id" uuid NOT NULL,
	"active" bool DEFAULT true,
	"created_at" timestamptz NOT NULL
);

ALTER TABLE users_company ADD PRIMARY KEY("user_id","company_id");

ALTER TABLE users_company
   ADD CONSTRAINT fk_users
   FOREIGN KEY ("user_id") 
   REFERENCES "users"("id");

ALTER TABLE users_company
   ADD CONSTRAINT fk_companies
   FOREIGN KEY ("company_id") 
   REFERENCES "companies"("id");

INSERT INTO "users_company"(user_id, company_id, active, created_at)
   VALUES ('` + value.USER_ID.String() + `', '` + value.COMPANY_ID.String() + `', true, CURRENT_TIMESTAMP)
   ON CONFLICT DO NOTHING;

CREATE TABLE "user_roles" (
	"id" uuid NOT NULL,
	"group_id" uuid NOT NULL,
	"path" varchar(255) NOT NULL,
	CONSTRAINT user_roles_pk PRIMARY KEY ("id")
);

CREATE TABLE "user_role_details" (
	"id" uuid NOT NULL,
	"user_role_id" uuid NOT NULL,
	"method" varchar(100) NOT NULL,
	CONSTRAINT user_role_details_pk PRIMARY KEY ("id")
);

ALTER TABLE user_roles
   ADD CONSTRAINT fk_user_groups
   FOREIGN KEY ("group_id") 
   REFERENCES "user_groups"("id");

ALTER TABLE user_role_details
   ADD CONSTRAINT fk_user_roles
   FOREIGN KEY ("user_role_id") 
   REFERENCES "user_roles"("id");

CREATE TABLE "item_categories" (
	"id" uuid NOT NULL,
	"company_id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT item_categories_pk PRIMARY KEY ("id")
);

ALTER TABLE item_categories
	ADD CONSTRAINT fk_companies
	FOREIGN KEY ("company_id") 
	REFERENCES "companies"("id");

CREATE TABLE "items" (
	"id" uuid NOT NULL,
	"company_id" uuid NOT NULL,
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

ALTER TABLE items
	ADD CONSTRAINT fk_companies
	FOREIGN KEY ("company_id") 
	REFERENCES "companies"("id");

CREATE TABLE "suppliers" (
	"id" uuid NOT NULL,
	"company_id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"phone" varchar(15) NULL,
	"address" text NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT suppliers_pk PRIMARY KEY ("id")
);

ALTER TABLE suppliers
	ADD CONSTRAINT fk_companies
	FOREIGN KEY ("company_id") 
	REFERENCES "companies"("id");

CREATE TABLE "sellers" (
	"id" uuid NOT NULL,
	"company_id" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"phone" varchar(15) NULL,
	"address" text NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT sellers_pk PRIMARY KEY ("id")
);

ALTER TABLE sellers
	ADD CONSTRAINT fk_companies
	FOREIGN KEY ("company_id") 
	REFERENCES "companies"("id");

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

ALTER TABLE transactions
	ADD CONSTRAINT fk_companies
	FOREIGN KEY ("company_id") 
	REFERENCES "companies"("id");

ALTER TABLE transactions
	ADD CONSTRAINT fk_users
	FOREIGN KEY ("user_id") 
	REFERENCES "users"("id");

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
	"company_id" uuid NOT NULL,
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

ALTER TABLE transaction_etc_types
	ADD CONSTRAINT fk_companies
	FOREIGN KEY ("company_id") 
	REFERENCES "companies"("id");
 
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
