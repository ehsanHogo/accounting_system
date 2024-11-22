CREATE TABLE "detaileds" (
  "id" BIGSERIAL PRIMARY KEY,
  "code" varchar UNIQUE,
  "title" varchar UNIQUE,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  "version" BIGINT DEFAULT 0 NOT NULL
);

CREATE TABLE subsidiaries (
  "id" BIGSERIAL PRIMARY KEY,
  "code" VARCHAR UNIQUE,
  "title" VARCHAR UNIQUE,
  "has_detailed" BOOLEAN,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  "version" BIGINT DEFAULT 0 NOT NULL
);

CREATE TABLE "vouchers" (
  "id" BIGSERIAL PRIMARY KEY,
  "number" varchar UNIQUE,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  "version" BIGINT DEFAULT 0 NOT NULL
);

CREATE TABLE "voucher_items" (
  "id" BIGSERIAL PRIMARY KEY,
  "voucher_id" BIGINT,
  "detailed_id" BIGINT,
  "subsidiary_id" BIGINT NOT NULL,
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
  FOREIGN KEY ("subsidiary_id") REFERENCES "subsidiaries" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;