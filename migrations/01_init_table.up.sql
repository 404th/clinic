CREATE TABLE IF NOT EXISTS "users" (
  "id" UUID DEFAULT (uuid_generate_v4()),
  "role_id" UUID,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "firstname" VARCHAR(255) NOT NULL,
  "surname" VARCHAR(255) NOT NULL,
  "email" VARCHAR NOT NULL,
  "wallet" BIGSERIAL NOT NULL DEFAULT 0,
  "active" INTEGER NOT NULL DEFAULT 1,
  "password" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW()),
  "deleted_at" TIMESTAMP,
  PRIMARY KEY ("id", "role_id")
);

CREATE TABLE IF NOT EXISTS "roles" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "rolename" VARCHAR(255) UNIQUE NOT NULL,
  "active" INTEGER NOT NULL DEFAULT 1,
  "is_paid" INTEGER NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW()),
  "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "queues" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "recipient_id" UUID NOT NULL,
  "customer_id" UUID NOT NULL,
  "paid_money" BIGSERIAL DEFAULT 0,
  "queue_number" INTEGER DEFAULT 1,
  "payment_status" INTEGER NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW()),
  "deleted_at" TIMESTAMP
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "queues" ADD FOREIGN KEY ("recipient_id") REFERENCES "users" ("id");
ALTER TABLE "queues" ADD FOREIGN KEY ("customer_id") REFERENCES "users" ("id");

-- COMMENT ON COLUMN IF NOT EXISTS "users"."wallet" IS 'USD';
-- COMMENT ON COLUMN IF NOT EXISTS "users"."active" IS '1-active, 0-inactive';
-- COMMENT ON COLUMN IF NOT EXISTS "roles"."active" IS '1-active, 0-inactive';
-- COMMENT ON COLUMN IF NOT EXISTS "roles"."is_paid" IS '1-paid, 0-unpaid';
-- COMMENT ON COLUMN IF NOT EXISTS "queues"."paid_money" IS 'USD';
-- COMMENT ON COLUMN IF NOT EXISTS "queues"."payment_status" IS '0-process, -1-cancelled, 1-paid';
