CREATE TABLE "detaileds" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE,
  "title" varchar UNIQUE,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE subsidiaries (
  "id" SERIAL PRIMARY KEY,
  "code" VARCHAR,
  "title" VARCHAR,
  "has_detailed" BOOLEAN,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "vouchers" (
  "id" SERIAL PRIMARY KEY,
  "number" varchar,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "voucher_items" (
  "id" SERIAL PRIMARY KEY,
  "voucher_id" INT,
  "detailed_id" INT,
  "subsidiary_id" INT,
  "debit" INT,
  "credit" INT,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

ALTER TABLE
  "voucher_items"
ADD
  FOREIGN KEY ("voucher_id") REFERENCES "vouchers" ("id");

-- ALTER TABLE "voucher_item" ADD FOREIGN KEY ("subsidiary_id") REFERENCES "subsidiary" ("id");
-- ALTER TABLE "voucher" ADD FOREIGN KEY ("voucher_items_id") REFERENCES "voucher_items" ("id");