-- migrate:up

CREATE TYPE "skytrax_type" AS ENUM ('Full Service', 'Low Cost');

CREATE DOMAIN "rating" AS SMALLINT NOT NULL CHECK (value BETWEEN 1 AND 5);

CREATE TABLE IF NOT EXISTS "airlines" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "skytrax_type" "skytrax_type" NOT NULL,
    "skytrax_rating" "rating" NOT NULL,
    "logo_id" BIGINT NULL DEFAULT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "airlines_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "airlines_logo_id_fkey" FOREIGN KEY ("logo_id") REFERENCES "images" ("id")
);

CREATE UNIQUE INDEX "airlines_name_unique"
ON "airlines" (UPPER("name"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "airlines";

DROP DOMAIN IF EXISTS "rating";

DROP TYPE IF EXISTS "skytrax_type";