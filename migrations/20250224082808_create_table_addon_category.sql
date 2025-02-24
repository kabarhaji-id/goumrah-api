-- migrate:up

CREATE TABLE IF NOT EXISTS "addon_categories" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "addon_categories_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "addon_categories_name_unique"
ON "addon_categories" (UPPER("name"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "addon_categories";