-- Remove foreign key constraints
-- ALTER TABLE "voucher" DROP CONSTRAINT IF EXISTS "voucher_voucher_items_id_fkey";
ALTER TABLE
    "voucher_items" DROP CONSTRAINT IF EXISTS "voucher_items_voucher_id_fkey";

ALTER TABLE
    "voucher_items" DROP CONSTRAINT IF EXISTS "voucher_item_subsidiary_id_fkey";

ALTER TABLE
    "voucher_items" DROP CONSTRAINT IF EXISTS "voucher_item_detailed_id_fkey";

-- Drop tables in reverse order of their dependencies
DROP TABLE IF EXISTS "voucher_items";

DROP TABLE IF EXISTS "vouchers";

DROP TABLE IF EXISTS "subsidiaries";

DROP TABLE IF EXISTS "detaileds";