-- migrate:up

ALTER TABLE "hotels"
ALTER COLUMN "description" TYPE VARCHAR(1000);

-- migrate:down

ALTER TABLE "hotels"
ALTER COLUMN "description" TYPE VARCHAR(500);