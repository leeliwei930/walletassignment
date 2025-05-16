-- Modify "wallets" table
ALTER TABLE "wallets" ALTER COLUMN "currency_code" TYPE character varying, ALTER COLUMN "currency_code" SET DEFAULT 'USD';
