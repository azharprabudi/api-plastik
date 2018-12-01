package execsql

// First ...
var First = `
CREATE TABLE "item_categories" (
	"id" serial NOT NULL,
	"name" varchar(100) NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT item_categories_pk PRIMARY KEY ("id")
);

CREATE TABLE "items" (
	"id" char(36) NOT NULL,
	"name" varchar(100) NOT NULL,
	"category_id" int NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY ("id")
);

ALTER TABLE items
   ADD CONSTRAINT fk_item_categories
   FOREIGN KEY ("category_id") 
   REFERENCES "item_categories"("id");

CREATE TABLE "suppliers" (
	"id" char(36) NOT NULL,
	"name" varchar(100) NOT NULL,
	"phone" varchar(15) NULL,
	"address" text NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT suppliers_pk PRIMARY KEY ("id")
);

CREATE TABLE "sellers" (
	"id" char(36) NOT NULL,
	"name" varchar(100) NOT NULL,
	"phone" varchar(15) NULL,
	"address" text NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT sellers_pk PRIMARY KEY ("id")
);

CREATE TABLE "expense_types" (
	"id" serial NOT NULL,
	"name" varchar(100) NOT NULL,
	CONSTRAINT expense_types_pk PRIMARY KEY ("id")
);

CREATE TABLE "expenses" (
	"id" serial NOT NULL,
	"expense_type_id" INT NOT NULL,
	"name" varchar(100) NOT NULL,
	CONSTRAINT expenses_pk PRIMARY KEY ("id")
);

ALTER TABLE expenses
   ADD CONSTRAINT fk_item_categories
   FOREIGN KEY ("expense_type_id") 
   REFERENCES "expense_types"("id");

`
