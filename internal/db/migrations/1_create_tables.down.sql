
ALTER TABLE
    "voucher_items" DROP CONSTRAINT IF EXISTS "voucher_items_voucher_id_fkey";

ALTER TABLE
    "voucher_items" DROP CONSTRAINT IF EXISTS "voucher_items_subsidiary_id_fkey";

ALTER TABLE
    "voucher_items" DROP CONSTRAINT IF EXISTS "voucher_items_detailed_id_fkey";


DROP TABLE IF EXISTS "voucher_items";

DROP TABLE IF EXISTS "vouchers";

DROP TABLE IF EXISTS "subsidiaries";

DROP TABLE IF EXISTS "detaileds";