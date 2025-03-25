-- migrate:up

ALTER TABLE "packages"
ADD COLUMN "fast_train" BOOLEAN DEFAULT FALSE;

-- migrate:down

ALTER TABLE "packages"
DROP COLUMN "fast_train";