-- migrate:up

ALTER TABLE "packages"
DROP COLUMN "description",
DROP COLUMN "is_active",
DROP COLUMN "is_recommended";

-- migrate:down

ALTER TABLE "packages"
ADD COLUMN "description" VARCHAR(500) NOT NULL DEFAULT '',
ADD COLUMN "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN "is_recommended" BOOLEAN NOT NULL DEFAULT FALSE;