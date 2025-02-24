-- migrate:up

CREATE TABLE IF NOT EXISTS "buses" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "seat" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "buses_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "buses_name_unique"
ON "buses" (UPPER("name"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "buses";