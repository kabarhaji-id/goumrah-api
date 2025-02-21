-- migrate:up

CREATE TABLE IF NOT EXISTS "images" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "src" VARCHAR(100) NOT NULL,
    "alt" VARCHAR(100) NOT NULL,
    "category" VARCHAR(100) NULL DEFAULT NULL,
    "title" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "images_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "images_src_unique"
ON "images" ("src")
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "images";
