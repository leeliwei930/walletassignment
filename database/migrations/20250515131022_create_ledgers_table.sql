-- Create "ledgers" table
CREATE TABLE "ledgers" (
    "id" uuid NOT NULL,
    "amount" bigint NOT NULL,
    "description" character varying NOT NULL,
    "transaction_type" character varying NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "wallet_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "ledgers_wallets_ledgers" FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "ledger_created_at" to table: "ledgers"
CREATE INDEX "ledger_created_at" ON "ledgers" ("created_at");
-- Create index "ledger_transaction_type" to table: "ledgers"
CREATE INDEX "ledger_transaction_type" ON "ledgers" ("transaction_type");
-- Create index "ledger_wallet_id_created_at" to table: "ledgers"
CREATE INDEX "ledger_wallet_id_created_at" ON "ledgers" ("wallet_id", "created_at");
-- Create index "ledger_wallet_id_transaction_type" to table: "ledgers"
CREATE INDEX "ledger_wallet_id_transaction_type" ON "ledgers" ("wallet_id", "transaction_type");
