-- migrate:up

ALTER TABLE "hotels"
ADD COLUMN "distance_landmark" VARCHAR(100) NOT NULL DEFAULT 'Masjidil Haram';

-- migrate:down

ALTER TABLE "hotels"
DROP COLUMN "distance_landmark";