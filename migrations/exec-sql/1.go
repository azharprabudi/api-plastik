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
	"unit_id" int NOT NULL,
	"created_at" timestamptz NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY ("id")
);

CREATE TABLE "item_units" (
	"id" serial NOT NULL,
	"name" varchar(100) NOT NULL,
	CONSTRAINT item_units_pk PRIMARY KEY ("id")
);

INSERT INTO "item_units"(id, name)
		VALUES (1, 'kg')
		ON CONFLICT DO NOTHING;

INSERT INTO "item_units"(id, name)
	VALUES (2, 'pcs')
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
	"id" char(36) NOT NULL,
	"expense_type_id" INT NOT NULL,
	"name" varchar(100) NOT NULL,
	"amount" numeric(20, 2) NOT NULL,
	"note" text null,
	CONSTRAINT expenses_pk PRIMARY KEY ("id")
);

ALTER TABLE expenses
   ADD CONSTRAINT fk_expense_type
   FOREIGN KEY ("expense_type_id") 
   REFERENCES "expense_types"("id");


CREATE TABLE "expense_images" (
	"id" serial NOT NULL,
	"image" varchar(100) NOT NULL,
	"expense_id" char(36) NOT NULL,
	CONSTRAINT expenses_pk PRIMARY KEY ("id")
);


ALTER TABLE expense_images
   ADD CONSTRAINT fk_expense
   FOREIGN KEY ("expense_id") 
   REFERENCES "expense_types"("id");

`
