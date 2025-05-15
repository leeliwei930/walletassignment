-- Create index "users_phone_number_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_number_key" ON "users" ("phone_number");
-- Create "wallets" table
CREATE TABLE "wallets" (
    "id" uuid NOT NULL,
    "balance" bigint NOT NULL DEFAULT 0,
    "currency_code" bigint NOT NULL DEFAULT 10,
    "decimal_places" bigint NOT NULL DEFAULT 2,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "user_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "wallets_users_wallets" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
