-- Create "users" table
CREATE TABLE "users" (
    "id" uuid NOT NULL,
    "first_name" character varying NOT NULL,
    "last_name" character varying NOT NULL,
    "phone_number" character varying NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("id")
);
-- Create index "user_phone_number" to table: "users"
CREATE UNIQUE INDEX "user_phone_number" ON "users" ("phone_number");
