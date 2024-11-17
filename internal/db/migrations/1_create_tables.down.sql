-- Remove foreign key constraints
ALTER TABLE "voucher" DROP CONSTRAINT IF EXISTS "voucher_voucher_items_id_fkey";
ALTER TABLE "voucher_item" DROP CONSTRAINT IF EXISTS "voucher_item_detailed_id_fkey";
ALTER TABLE "voucher_item" DROP CONSTRAINT IF EXISTS "voucher_item_subsidiary_id_fkey";

-- Drop tables in reverse order of their dependencies
DROP TABLE IF EXISTS "voucher_item";
DROP TABLE IF EXISTS "voucher_items";
DROP TABLE IF EXISTS "voucher";
DROP TABLE IF EXISTS "subsidiary";
DROP TABLE IF EXISTS "detailed";
