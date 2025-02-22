-- migrate:up

CREATE TYPE "package_category" AS ENUM ('Silver', 'Gold', 'Platinum', 'Luxury');

CREATE TYPE "package_type" AS ENUM ('Reguler', 'Plus');

CREATE TABLE IF NOT EXISTS "packages" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "thumbnail_id" BIGINT NULL DEFAULT NULL,
    "name" VARCHAR(100) NOT NULL,
    "description" VARCHAR(500) NOT NULL,
    "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
    "category" "package_category" NOT NULL,
    "type" "package_type" NOT NULL,
    "slug" VARCHAR(105) NOT NULL,
    "is_recommended" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "packages_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "packages_thumbnail_id_fkey" FOREIGN KEY ("thumbnail_id") REFERENCES "images" ("id")
);

CREATE UNIQUE INDEX "packages_name_unique"
ON "packages" (UPPER("name"))
WHERE "deleted_at" IS NULL;

CREATE UNIQUE INDEX "packages_slug_unique"
ON "packages" (UPPER("slug"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "packages";

DROP TYPE IF EXISTS "package_category";

DROP TYPE IF EXISTS "package_type";
