-- migrate:up

CREATE TABLE IF NOT EXISTS "hotels" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "rating" "rating" NOT NULL,
    "map" TEXT NOT NULL,
    "address" VARCHAR(500) NOT NULL,
    "distance" DECIMAL(6, 2) NOT NULL,
    "review" TEXT NOT NULL,
    "description" VARCHAR(500) NOT NULL,
    "location" VARCHAR(100) NOT NULL,
    "slug" VARCHAR(105) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "hotels_id_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "hotels_name_unique"
ON "hotels" (UPPER("name"))
WHERE "deleted_at" IS NULL;

CREATE UNIQUE INDEX "hotels_slug_unique"
ON "hotels" (UPPER("slug"))
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP TABLE IF EXISTS "hotels";
