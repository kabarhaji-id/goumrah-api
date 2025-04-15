-- migrate:up

CREATE TABLE IF NOT EXISTS "landing_travel_destination_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_travel_destination_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_travel_destination_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_travel_destination_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_travel_destination_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_travel_destination_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing travel destination content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_travel_destination_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_travel_destination_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing travel destination content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_travel_destination_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_travel_destination_content_exists
BEFORE UPDATE ON landing_section_headers
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_travel_destination_content_exists();

CREATE TRIGGER prevent_insert_landing_travel_destination_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_travel_destination_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_travel_destination_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_travel_destination_content_destinations" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE,
    "image_id" BIGINT NULL DEFAULT NULL,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_travel_destination_content_destinations_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_travel_destination_content_destinations_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION set_travel_destination_content_destinations_image_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_travel_destination_content_destinations SET image_id = NULL WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_travel_destination_content_destinations_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.image_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing travel destination content destination with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_travel_destination_content_destinations_image_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_travel_destination_content_destinations_image_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_travel_destination_content_destinations_image_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_landing_travel_destination_content_destinations_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_travel_destination_content_destinations
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_travel_destination_content_destinations_if_image_is_soft_deleted();

INSERT INTO "landing_section_headers" (
    "title", "subtitle", "tags_line"
) VALUES
    ('Temukan Destinasi Wisata Favorit', '#EpicMoment', NULL);

INSERT INTO "landing_travel_destination_content" (
    "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id"
) VALUES
    (1, TRUE, TRUE, TRUE, 10);

INSERT INTO "landing_travel_destination_content_destinations" (
    "is_enabled", "is_mobile", "is_desktop", "image_id", "name"
) VALUES
    (TRUE, TRUE, TRUE, NULL, 'Al-Ula'),
    (TRUE, TRUE, TRUE, NULL, 'Thaif'),
    (TRUE, TRUE, TRUE, NULL, 'Dubai');

-- migrate:down

DROP TRIGGER IF EXISTS prevent_insert_landing_travel_destination_content_destinations_if_image_is_soft_deleted ON landing_travel_destination_content_destinations;
DROP TRIGGER IF EXISTS set_travel_destination_content_destinations_image_id_null_on_image_soft_deleted ON images;
DROP FUNCTION IF EXISTS prevent_insert_landing_travel_destination_content_destinations_if_image_is_soft_deleted();
DROP FUNCTION IF EXISTS set_travel_destination_content_destinations_image_id_null_on_image_soft_deleted();

DROP TABLE IF EXISTS "landing_travel_destination_content_destinations";

DROP TRIGGER IF EXISTS prevent_insert_landing_travel_destination_content_if_landing_section_header_is_soft_deleted ON landing_travel_destination_content;
DROP TRIGGER IF EXISTS prevent_delete_landing_section_header_if_landing_travel_destination_content_exists ON landing_section_headers;
DROP FUNCTION IF EXISTS prevent_insert_landing_travel_destination_content_if_landing_section_header_is_soft_deleted();
DROP FUNCTION IF EXISTS prevent_delete_landing_section_header_if_landing_travel_destination_content_exists();

DROP TABLE IF EXISTS "landing_travel_destination_content";
