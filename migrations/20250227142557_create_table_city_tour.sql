-- migrate:up

CREATE TABLE IF NOT EXISTS "city_tours" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "city" VARCHAR(100) NOT NULL,
    "description" VARCHAR(500) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "city_tours_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "city_tours_name_unique"
ON "city_tours" (UPPER("name"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "city_tours";