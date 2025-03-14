-- migrate:up

CREATE TABLE IF NOT EXISTS "hotel_images" (
    "hotel_id" BIGINT NOT NULL,
    "image_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "hotel_images_id_pkey" PRIMARY KEY ("hotel_id", "image_id"),
    CONSTRAINT "hotel_images_hotel_id_fkey" FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id"),
    CONSTRAINT "hotel_images_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION delete_hotel_image_on_hotel_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE hotel_images SET deleted_at = NOW() WHERE hotel_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_hotel_image_if_hotel_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM hotels WHERE id = NEW.hotel_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert hotel image with soft deleted hotel'
            USING ERRCODE = '23503', CONSTRAINT = 'hotel_images_hotel_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_hotel_image_on_hotel_soft_deleted
BEFORE UPDATE ON hotels
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_hotel_image_on_hotel_soft_deleted();

CREATE TRIGGER prevent_insert_hotel_image_if_hotel_is_soft_deleted
BEFORE INSERT OR UPDATE ON hotel_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_hotel_image_if_hotel_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_hotel_image_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE hotel_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_hotel_image_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert hotel image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'hotel_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_hotel_image_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_hotel_image_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_hotel_image_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON hotel_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_hotel_image_if_image_is_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS prevent_insert_hotel_image_if_image_is_soft_deleted ON hotel_images;
DROP TRIGGER IF EXISTS delete_hotel_image_on_image_soft_deleted ON images;
DROP FUNCTION IF EXISTS prevent_insert_hotel_image_if_image_is_soft_deleted;
DROP FUNCTION IF EXISTS delete_hotel_image_on_image_soft_deleted;

DROP TRIGGER IF EXISTS prevent_insert_hotel_image_if_hotel_is_soft_deleted ON hotel_images;
DROP TRIGGER IF EXISTS delete_hotel_image_on_hotel_soft_deleted ON hotels;
DROP FUNCTION IF EXISTS prevent_insert_hotel_image_if_hotel_is_soft_deleted;
DROP FUNCTION IF EXISTS delete_hotel_image_on_hotel_soft_deleted;

DROP TABLE IF EXISTS "hotel_images";
