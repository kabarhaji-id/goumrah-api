-- migrate:up

CREATE TABLE IF NOT EXISTS "airports" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "city" VARCHAR(100) NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "code" CHAR(3) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "airports_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "airports_name_unique"
ON "airports" (UPPER("name"))
WHERE "deleted_at" IS NULL;

CREATE UNIQUE INDEX "airports_code_unique"
ON "airports" (UPPER("code"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "airports";