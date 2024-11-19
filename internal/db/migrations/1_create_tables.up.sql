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
  "code" VARCHAR UNIQUE,
  "title" VARCHAR UNIQUE,
  "has_detailed" BOOLEAN,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "vouchers" (
  "id" SERIAL PRIMARY KEY,
  "number" varchar UNIQUE,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "voucher_items" (
  "id" SERIAL PRIMARY KEY,
  "voucher_id" INT,
  "detailed_id" INT NOT NULL,
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
  FOREIGN KEY ("voucher_id") REFERENCES "vouchers" ("id") ON DELETE CASCADE;

ALTER TABLE
  "voucher_items"
ADD
  FOREIGN KEY ("detailed_id") REFERENCES "detaileds" ("id") ON DELETE RESTRICT;

ALTER TABLE
  "voucher_items"
ADD
  FOREIGN KEY ("subsidiary_id") REFERENCES "subsidiaries" ("id");

-- ALTER TABLE "voucher_item" ADD FOREIGN KEY ("subsidiary_id") REFERENCES "subsidiary" ("id");
-- ALTER TABLE "voucher" ADD FOREIGN KEY ("voucher_items_id") REFERENCES "voucher_items" ("id");