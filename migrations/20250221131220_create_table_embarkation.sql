-- migrate:up

CREATE TABLE IF NOT EXISTS "embarkations" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "latitude" FLOAT NOT NULL,
    "longitude" FLOAT NOT NULL,
    "slug" VARCHAR(105) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "embarkations_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "embarkations_name_unique"
ON "embarkations" (UPPER("name"))
WHERE "deleted_at" IS NULL;

CREATE UNIQUE INDEX "embarkations_slug_unique"
ON "embarkations" (UPPER("slug"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "embarkations";