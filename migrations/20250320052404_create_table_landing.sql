-- migrate:up

CREATE TABLE IF NOT EXISTS "landing_hero_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "title" VARCHAR(100) NOT NULL,
    "description" VARCHAR(500) NOT NULL,
    "tags_line" VARCHAR(50) NOT NULL,
    "button_label" VARCHAR(100) NOT NULL,
    "button_url" TEXT NULL DEFAULT NULL,
    "image_id" INT NULL DEFAULT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_hero_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_hero_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_hero_content_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION set_landing_hero_content_image_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_hero_content SET image_id = NULL WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_hero_content_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.image_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing hero content with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_hero_content_image_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_landing_hero_content_image_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_landing_hero_content_image_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_landing_hero_content_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_hero_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_hero_content_if_image_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_section_headers" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "title" VARCHAR(100) NOT NULL,
    "subtitle" VARCHAR(100) NULL DEFAULT NULL,
    "tags_line" VARCHAR(50) NULL DEFAULT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_section_headers_id_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "landing_package_items" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "package_id" BIGINT NOT NULL,
    "button_label" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_package_items_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_package_items_package_id_fkey" FOREIGN KEY ("package_id") REFERENCES "packages" ("id")
);

CREATE OR REPLACE FUNCTION delete_landing_package_item_on_package_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_package_items SET deleted_at = NOW() WHERE package_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_package_item_if_package_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM packages WHERE id = NEW.package_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package item with soft deleted package'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_items_package_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_landing_package_item_on_package_soft_deleted
BEFORE UPDATE ON packages
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_landing_package_item_on_package_soft_deleted();

CREATE TRIGGER prevent_insert_landing_package_item_if_package_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_package_items
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_package_item_if_package_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_single_package_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "silver_landing_package_item_id" BIGINT NULL DEFAULT NULL,
    "gold_landing_package_item_id" BIGINT NULL DEFAULT NULL,
    "platinum_landing_package_item_id" BIGINT NULL DEFAULT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_single_package_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_single_package_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_single_package_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id"),
    CONSTRAINT "landing_single_package_content_silver_landing_package_item_id_fkey" FOREIGN KEY ("silver_landing_package_item_id") REFERENCES "landing_package_items" ("id"),
    CONSTRAINT "landing_single_package_content_gold_landing_package_item_id_fkey" FOREIGN KEY ("gold_landing_package_item_id") REFERENCES "landing_package_items" ("id"),
    CONSTRAINT "landing_single_package_content_platinum_landing_package_item_id_fkey" FOREIGN KEY ("platinum_landing_package_item_id") REFERENCES "landing_package_items" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_single_package_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_single_package WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing single package content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_single_package_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION set_landing_single_package_content_landing_package_item_id_null_on_landing_package_item_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NOT NULL THEN
        UPDATE landing_single_package_content SET silver_landing_package_item_id = NULL WHERE silver_landing_package_item_id = OLD.id;
        UPDATE landing_single_package_content SET gold_landing_package_item_id = NULL WHERE gold_landing_package_item_id = OLD.id;
        UPDATE landing_single_package_content SET platinum_landing_package_item_id = NULL WHERE platinum_landing_package_item_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_single_package_content_if_landing_package_item_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.silver_landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted silver landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_silver_landing_package_item_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.gold_landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted gold landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_gold_landing_package_item_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.platinum_landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted platinum landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_platinum_landing_package_item_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_single_package_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_single_package_content_exists();

CREATE TRIGGER prevent_insert_landing_single_package_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_single_package_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_single_package_content_if_landing_section_header_is_soft_deleted();

CREATE TRIGGER set_landing_single_package_content_landing_package_item_id_null_on_landing_package_item_soft_deleted
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_landing_single_package_content_landing_package_item_id_null_on_landing_package_item_soft_deleted();

CREATE TRIGGER prevent_insert_landing_single_package_content_if_landing_package_item_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_single_package_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_single_package_content_if_landing_package_item_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_package_details" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_package_details_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_package_details_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_package_detail_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_package_details WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_details_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_package_detail_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package detail with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_details_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_package_detail_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_package_detail_exists();

CREATE TRIGGER prevent_insert_landing_package_detail_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_package_details
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_package_detail_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_package_detail_items" (
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_package_detail_id" BIGINT NOT NULL,
    "landing_package_item_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_package_detail_items_pkey" PRIMARY KEY ("landing_package_detail_id", "landing_package_item_id"),
    CONSTRAINT "landing_package_detail_items_landing_package_detail_id_fkey" FOREIGN KEY ("landing_package_detail_id") REFERENCES "landing_package_details" ("id"),
    CONSTRAINT "landing_package_detail_items_landing_package_item_id_fkey" FOREIGN KEY ("landing_package_item_id") REFERENCES "landing_package_items" ("id")
);

CREATE OR REPLACE FUNCTION delete_landing_package_detail_item_on_landing_package_detail_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_package_detail_items SET deleted_at = NOW() WHERE landing_package_detail_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_package_detail_item_if_landing_package_detail_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package detail item with soft deleted landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_detail_items_landing_package_detail_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_landing_package_detail_item_on_landing_package_detail_soft_deleted
BEFORE UPDATE ON landing_package_details
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_landing_package_detail_item_on_landing_package_detail_soft_deleted();

CREATE TRIGGER prevent_insert_landing_package_detail_item_if_landing_package_detail_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_package_detail_items
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_package_detail_item_if_landing_package_detail_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_landing_package_detail_item_on_landing_package_item_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_package_detail_items SET deleted_at = NOW() WHERE landing_package_item_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_package_detail_item_if_landing_package_item_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package detail item with soft deleted landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_detail_items_landing_package_item_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_landing_package_detail_item_on_landing_package_item_soft_deleted
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_landing_package_detail_item_on_landing_package_item_soft_deleted();

CREATE TRIGGER prevent_insert_landing_package_detail_item_if_landing_package_item_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_package_detail_items
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_package_detail_item_if_landing_package_item_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_packages_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "silver_landing_package_detail_id" BIGINT NOT NULL,
    "gold_landing_package_detail_id" BIGINT NOT NULL,
    "platinum_landing_package_detail_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_packages_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_packages_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_packages_content_silver_landing_package_detail_id_fkey" FOREIGN KEY ("silver_landing_package_detail_id") REFERENCES "landing_package_details" ("id"),
    CONSTRAINT "landing_packages_content_gold_landing_package_detail_id_fkey" FOREIGN KEY ("gold_landing_package_detail_id") REFERENCES "landing_package_details" ("id"),
    CONSTRAINT "landing_packages_content_platinum_landing_package_detail_id_fkey" FOREIGN KEY ("platinum_landing_package_detail_id") REFERENCES "landing_package_details" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_package_detail_if_landing_packages_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_packages_content WHERE silver_landing_package_detail_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing package detail with existing landing packages content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_silver_landing_package_detail_id_fkey';
    END IF;
    IF EXISTS (SELECT 1 FROM landing_packages_content WHERE gold_landing_package_detail_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing package detail with existing landing packages content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_gold_landing_package_detail_id_fkey';
    END IF;
    IF EXISTS (SELECT 1 FROM landing_packages_content WHERE platinum_landing_package_detail_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing package detail with existing landing packages content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_platinum_landing_package_detail_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_packages_content_if_landing_package_detail_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.silver_landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing packages content with soft deleted silver landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_silver_landing_package_detail_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.gold_landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing packages content with soft deleted gold landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_gold_landing_package_detail_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.platinum_landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing packages content with soft deleted platinum landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_platinum_landing_package_detail_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_package_detail_if_landing_packages_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_package_detail_if_landing_packages_content_exists();

CREATE TRIGGER prevent_insert_landing_packages_content_if_landing_package_detail_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_packages_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_packages_content_if_landing_package_detail_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_features_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "footer_title" VARCHAR(100) NOT NULL,
    "button_about" VARCHAR(100) NOT NULL,
    "button_package" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_features_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_features_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_features_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_features_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_features_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing features content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_features_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_features_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing features content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_features_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_features_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_features_content_exists();

CREATE TRIGGER prevent_insert_landing_features_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_features_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_features_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_features_content_benefits" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "title" VARCHAR(100) NOT NULL,
    "subtitle" VARCHAR(500) NOT NULL,
    "logo_id" BIGINT NULL DEFAULT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_features_content_benefits_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_features_content_benefits_logo_id_fkey" FOREIGN KEY ("logo_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION set_landing_features_content_benefits_logo_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_features_content_benefits SET logo_id = NULL WHERE logo_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_features_content_benefits_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.logo_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.logo_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing features content benefits with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_features_content_benefits_logo_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_landing_features_content_benefits_logo_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_landing_features_content_benefits_logo_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_landing_features_content_benefits_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_features_content_benefits
FOR EACH ROW
WHEN (NEW.logo_id IS NOT NULL)
EXECUTE FUNCTION prevent_insert_landing_features_content_benefits_if_image_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_moments_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_moments_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_moments_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_moments_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_moments_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_moments_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing moments content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_moments_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_moments_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing moments content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_moments_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_moments_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_moments_content_exists();

CREATE TRIGGER prevent_insert_landing_moments_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_moments_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_moments_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_moments_content_images" (
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "image_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_moments_content_images_pkey" PRIMARY KEY ("image_id"),
    CONSTRAINT "landing_moments_content_images_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION delete_landing_moments_content_image_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        DELETE FROM landing_moments_content_images WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_moments_content_image_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing moments content image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_moments_content_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_landing_moments_content_image_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_landing_moments_content_image_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_landing_moments_content_image_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_moments_content_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_moments_content_image_if_image_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_affiliates_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_affiliates_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_affiliates_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_affiliates_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_affiliates_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_affiliates_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing affiliates content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_affiliates_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_affiliates_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing affiliates content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_affiliates_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_affiliates_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_affiliates_content_exists();

CREATE TRIGGER prevent_insert_landing_affiliates_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_affiliates_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_affiliates_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_affiliates_content_affiliates" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "name" VARCHAR(100) NOT NULL,
    "logo_id" BIGINT NULL DEFAULT NULL,
    "width" INT NOT NULL,
    "height" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_affiliates_content_affiliates_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_affiliates_content_affiliates_logo_id_fkey" FOREIGN KEY ("logo_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION set_landing_affiliates_content_affiliates_logo_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_affiliates_content_affiliates SET logo_id = NULL WHERE logo_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_affiliates_content_affiliates_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.logo_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.logo_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing affiliates content affiliates with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_affiliates_content_affiliates_logo_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_landing_affiliates_content_affiliates_logo_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_landing_affiliates_content_affiliates_logo_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_landing_affiliates_content_affiliates_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_affiliates_content_affiliates
FOR EACH ROW
WHEN (NEW.logo_id IS NOT NULL)
EXECUTE FUNCTION prevent_insert_landing_affiliates_content_affiliates_if_image_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_faq_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_faq_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_faq_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_faq_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_faq_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_faq_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing faq content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_faq_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_faq_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing faq content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_faq_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_faq_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_faq_content_exists();

CREATE TRIGGER prevent_insert_landing_faq_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_faq_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_faq_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_faq_content_faqs" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "question" VARCHAR(100) NOT NULL,
    "answer" VARCHAR(500) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_faq_content_faqs_id_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "landing_menus" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "icon" VARCHAR(100) NOT NULL,
    "label" VARCHAR(100) NOT NULL,
    "path" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_menus_id_pkey" PRIMARY KEY ("id")
);

INSERT INTO "landing_hero_content" (
    "title", "description", "tags_line", "button_label", "button_url", "image_id"
) VALUES (
    'Rumah ke Makkah, Hanya Satu Langkah',
    'Wujudkan perjalanan Umrah impian anda dengan mudah dan terpercaya, dari mana saja, kapan saja.',
    '#bikinTenang',
    'Cek Izin Umroh Kami Disini',
    'https://simpu.kemenag.go.id/home/detail/3039',
    NULL
);

INSERT INTO "landing_section_headers" (
    "title", "subtitle", "tags_line"
) VALUES
    ('Umrah Ideal dengan Momen Tak Terlupakan mulai dari 22 jt', 'Sambut Panggilan-Nya', NULL),
    ('Paket Rekomendasi Silver', 'Paket Hemat, Ibadah Khidmat', NULL),
    ('Paket Rekomendasi Gold', 'Pilihan Bijak Untuk Perjalanan Penuh Makna', NULL),
    ('Paket Rekomendasi Platinum', 'Ibadah Tenang, Nyaman Maksimal', NULL),
    ('Berangkat Umroh Bersama goumrah.id', 'Dapatkan Kelebihannya', NULL),
    ('Abadikan Moment Tak Terlupakan Bersama goumrah.id', '#EpicMoment', NULL),
    ('Afiliasi Kami', '#EpicMoment', NULL),
    ('Testimoni jama''ah', 'Apa yang para Jama''ah katakan tentang kami', NULL);

INSERT INTO "landing_single_package_content" (
    "landing_section_header_id", "silver_landing_package_item_id", "gold_landing_package_item_id", "platinum_landing_package_item_id"
) VALUES (1, NULL, NULL, NULL);

INSERT INTO "landing_package_details" (
    "landing_section_header_id"
) VALUES (2), (3), (4);

INSERT INTO "landing_packages_content" (
    "silver_landing_package_detail_id", "gold_landing_package_detail_id", "platinum_landing_package_detail_id"
) VALUES (1, 2, 3);

INSERT INTO "landing_features_content" (
    "landing_section_header_id", "footer_title", "button_about", "button_package"
) VALUES (
    5,
    'Jadi, Tunggu apa lagi? Percayakan perjalanan Ibadah Umroh kamu bersama goumrah.id',
    'Kenalan Yuk sama goumrah.id',
    'Lihat Semua Paket Umroh'
);

INSERT INTO "landing_features_content_benefits" (
    "title", "subtitle", "logo_id"
) VALUES
    ('Pelayanan Terbaik', 'goumrah.id memberikan pelayanan pemesanan paket umroh dengan fasilitas terbaik untuk setiap jama''ah', NULL),
    ('Pemesanan Mudah', 'Kami berkomitmen memberikan kemudahan untuk setiap pemesanan paket umroh jama''ah dari awal hingga akhir.', NULL),
    ('Biaya Transparan', 'Sering mendapatkan biaya tak terduga saat pesan paket umroh? disini tidak lagi! kami memberikan laporan biaya yang transparan.', NULL),
    ('Terpercaya', 'goumrah.id sudah berpengalaman 18 tahun memberangkatkan jama''ah ke Tanah Suci dan sudah berizin PPIU di Kemenag.', NULL),
    ('Pembayaran Aman', 'goumrah.id sudah bekerjasama dengan platform pembayaran terpercaya di Indonesia, dan sudah berizin di OJK ataupun Bank Indonesia.', NULL);

INSERT INTO "landing_moments_content" (
    "landing_section_header_id"
) VALUES (6);

INSERT INTO "landing_affiliates_content" (
    "landing_section_header_id"
) VALUES (7);

INSERT INTO "landing_affiliates_content_affiliates" (
    "name", "logo_id", "width", "height"
) VALUES
    ('Kementerian Agama', NULL, 84, 66),
    ('Sistem Pengawasan Umrah', NULL, 84, 66),
    ('Komite Akreditasi Nasional', NULL, 84, 66),
    ('5 Pasti', NULL, 84, 66),
    ('Himpunan Penyelenggara Umrah dan Haji', NULL, 84, 66),
    ('Association of The Indonesian Tours and Travel Agencies', NULL, 84, 66),
    ('Badan Nasional Sertifikasi Profesi', NULL, 84, 66),
    ('Garuda Indonesia', NULL, 84, 66),
    ('Lion Air', NULL, 84, 66);

INSERT INTO "landing_faq_content" (
    "landing_section_header_id"
) VALUES (8);

INSERT INTO "landing_faq_content_faqs" (
    "question", "answer"
) VALUES
    ('Apa itu goumrah.id?', 'Goumrah.id adalah biro perjalanan umroh yang **berizin resmi dari Kementerian Agama RI**. Kami menyediakan berbagai pilihan paket umroh, mulai dari yang hemat hingga eksklusif, untuk memenuhi kebutuhan perjalanan ibadah kamu.'),
    ('Apakah goumrah.id memiliki izin resmi dari Kementerian Agama?', '**Ya, goumrah.id memiliki izin resmi** sebagai **Penyelenggara Perjalanan Ibadah Umroh (PPIU)** dari Kementerian Agama RI. Nomor izin kami adalah **No.27052200387740007.**'),
    ('Dimana lokasi kantor pusat goumrah.id?', 'Kantor pusat kami berlokasi di **JL. GM Ainul Yakin Blk. A-B No.35, Kalibata, Kec. Pancoran, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12740**, atau bisa [lihat disini.](https://maps.app.goo.gl/fMW2ZttUt9wJcCSz6)');

INSERT INTO "landing_menus" (
    "icon", "label", "path"
) VALUES
    ('HomeIcon', 'Beranda', '/'),
    ('KaabaIcon', 'Paket', '/paket'),
    ('BlogIcon', 'Blog', '/blog'),
    ('FaqIcon', 'FAQ', '/faq'),
    ('AboutIcon', 'About', '/about');

-- migrate:down

DROP TABLE IF EXISTS "landing_menus";

DROP TABLE IF EXISTS "landing_faq_content_faqs";

DROP TABLE IF EXISTS "landing_faq_content";

DROP TRIGGER IF EXISTS prevent_insert_landing_affiliates_content_affiliates_if_image_is_soft_deleted ON landing_affiliates_content_affiliates;
DROP TRIGGER IF EXISTS set_landing_affiliates_content_affiliates_logo_id_null_on_image_soft_deleted ON images;
DROP TABLE IF EXISTS "landing_affiliates_content_affiliates";

DROP TABLE IF EXISTS "landing_affiliates_content";

DROP TRIGGER IF EXISTS prevent_insert_landing_moments_content_image_if_image_is_soft_deleted ON landing_moments_content_images;
DROP TRIGGER IF EXISTS delete_landing_moments_content_image_on_image_soft_deleted ON images;
DROP TABLE IF EXISTS "landing_moments_content_images";

DROP TABLE IF EXISTS "landing_moments_content";

DROP TRIGGER IF EXISTS prevent_insert_landing_features_content_benefits_if_image_is_soft_deleted ON landing_features_content_benefits;
DROP TRIGGER IF EXISTS set_landing_features_content_benefits_logo_id_null_on_image_soft_deleted ON images;
DROP TABLE IF EXISTS "landing_features_content_benefits";

DROP TABLE IF EXISTS "landing_features_content";

DROP TRIGGER IF EXISTS prevent_insert_landing_packages_content_if_landing_package_detail_is_soft_deleted ON landing_packages;
DROP TRIGGER IF EXISTS prevent_delete_landing_package_detail_if_landing_packages_content_exists ON landing_package_items;
DROP TABLE IF EXISTS "landing_packages_content";

DROP TRIGGER IF EXISTS prevent_insert_landing_package_detail_item_if_landing_package_item_is_soft_deleted ON landing_package_detail_items;
DROP TRIGGER IF EXISTS delete_landing_package_detail_item_on_landing_package_item_soft_deleted ON landing_package_items;
DROP TRIGGER IF EXISTS prevent_insert_landing_package_detail_item_if_landing_package_detail_is_soft_deleted ON landing_package_detail_items;
DROP TRIGGER IF EXISTS delete_landing_package_detail_item_on_landing_package_detail_soft_deleted ON landing_package_details;
DROP TABLE IF EXISTS "landing_package_detail_items";

DROP TRIGGER IF EXISTS prevent_insert_landing_package_detail_if_landing_section_header_is_soft_deleted ON landing_package_details;
DROP TRIGGER IF EXISTS prevent_delete_landing_section_header_if_landing_package_detail_exists ON landing_package_items;
DROP TABLE IF EXISTS "landing_package_details";

DROP TRIGGER IF EXISTS prevent_insert_landing_single_package_content_if_landing_package_item_is_soft_deleted ON landing_single_package_content;
DROP TRIGGER IF EXISTS set_landing_single_package_content_landing_package_item_id_null_on_landing_package_item_soft_deleted ON landing_package_items;
DROP TRIGGER IF EXISTS prevent_insert_landing_single_package_content_if_landing_section_header_is_soft_deleted ON landing_single_package_content;
DROP TRIGGER IF EXISTS prevent_delete_landing_section_header_if_landing_single_package_content_exists ON landing_package_items;
DROP TABLE IF EXISTS "landing_single_package_content";

DROP TRIGGER IF EXISTS prevent_insert_landing_package_item_if_package_is_soft_deleted ON landing_package_items;
DROP TRIGGER IF EXISTS delete_landing_package_item_on_package_soft_deleted ON packages;
DROP TABLE IF EXISTS "landing_package_items";

DROP TABLE IF EXISTS "landing_section_headers";

DROP TRIGGER IF EXISTS prevent_insert_landing_hero_content_if_image_is_soft_deleted ON landing_hero_content;
DROP TRIGGER IF EXISTS set_landing_hero_content_image_id_null_on_image_soft_deleted ON images;
DROP TABLE IF EXISTS "landing_hero_content";