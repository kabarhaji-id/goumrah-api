-- migrate:up

CREATE TABLE IF NOT EXISTS "facilities" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "icon" VARCHAR(100) NOT NULL,
    
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "facilities_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "facilities_name_unique"
ON "facilities" (UPPER("name"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "facilities";
