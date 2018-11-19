package execsql

// First ...
var First = `
CREATE TABLE "itemCategory" (
	"itemCategoryId" serial NOT NULL,
	"itemCategoryName" varchar(100) NOT NULL,
	"createdAt" timestamptz NOT NULL,
	CONSTRAINT item_category_pk PRIMARY KEY ("itemCategoryId")
);

CREATE TABLE "item" (
	"itemId" char(36) NOT NULL,
	"itemName" varchar(100) NOT NULL,
	"itemCategoryId" int NOT NULL,
	"createdAt" timestamptz NOT NULL,
	CONSTRAINT item_pk PRIMARY KEY ("itemId")
);

ALTER TABLE item 
   ADD CONSTRAINT fk_item_category
   FOREIGN KEY ("itemCategoryId") 
   REFERENCES "itemCategory"("itemCategoryId");

CREATE TABLE "supplier" (
	"supplierId" char(36) NOT NULL,
	"supplierName" varchar(100) NOT NULL,
	"supplierPhone" varchar(15) NULL,
	"supplierAddress" text NULL,
	"createdAt" timestamptz NOT NULL,
	CONSTRAINT supplier_pk PRIMARY KEY ("supplierId")
);

CREATE TABLE "seller" (
	"sellerId" char(36) NOT NULL,
	"sellerName" varchar(100) NOT NULL,
	"sellerPhone" varchar(15) NULL,
	"sellerAddress" text NULL,
	"createdAt" timestamptz NOT NULL,
	CONSTRAINT seller_pk PRIMARY KEY ("sellerId")
);
`
