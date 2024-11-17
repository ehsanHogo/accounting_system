CREATE TABLE "detailed" (
  "id" integer PRIMARY KEY,
  "code" varchar,
  "title" varchar
);

CREATE TABLE "subsidiary" (
  "id" integer PRIMARY KEY,
  "code" varchar,
  "title" varchar,
  "has_detailed" bool
);

CREATE TABLE "voucher" (
  "id" integer PRIMARY KEY,
  "number" varchar,
  "voucher_items_id" integer
);

CREATE TABLE "voucher_items" (
  "id" integer PRIMARY KEY,
  "voucher_item_id" integer
);

CREATE TABLE "voucher_item" (
  "id" integer PRIMARY KEY,
  "detailed_id" integer,
  "subsidiary_id" integer,
  "debit" integer,
  "credit" integer
);

ALTER TABLE "voucher_item" ADD FOREIGN KEY ("detailed_id") REFERENCES "detailed" ("id");

ALTER TABLE "voucher_item" ADD FOREIGN KEY ("subsidiary_id") REFERENCES "subsidiary" ("id");

ALTER TABLE "voucher" ADD FOREIGN KEY ("voucher_items_id") REFERENCES "voucher_items" ("id");