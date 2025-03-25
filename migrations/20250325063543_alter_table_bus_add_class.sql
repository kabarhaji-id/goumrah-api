-- migrate:up

DROP INDEX IF EXISTS "buses_name_unique";

CREATE TYPE "bus_class" AS ENUM ('Economy', 'VIP');

ALTER TABLE "buses"
ADD COLUMN "class" "bus_class" NOT NULL DEFAULT 'Economy';

CREATE UNIQUE INDEX "buses_name_class_unique"
ON "buses" (UPPER("name"), "class")
WHERE "deleted_at" IS NULL;

-- migrate:down

DROP INDEX IF EXISTS "buses_name_class_unique";

ALTER TABLE "buses"
DROP COLUMN "class";

DROP TYPE IF EXISTS "bus_class";

CREATE UNIQUE INDEX "buses_name_unique"
ON "buses" (UPPER("name"))
WHERE "deleted_at" IS NULL;