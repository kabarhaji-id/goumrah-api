-- migrate:up

CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "full_name" VARCHAR(100) NOT NULL,
    "phone_number" VARCHAR(20) NOT NULL,
    "email" VARCHAR(256) NOT NULL,
    "address" VARCHAR(500) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "users_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_phone_number_unique"
ON "users" (UPPER("phone_number"))
WHERE "deleted_at" IS NULL;

CREATE UNIQUE INDEX "users_email_unique"
ON "users" (UPPER("email"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "users";