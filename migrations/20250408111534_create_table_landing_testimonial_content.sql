-- migrate:up

CREATE DOMAIN "decimal_rating" AS DECIMAL NOT NULL CHECK (value BETWEEN 1 AND 5);

CREATE TABLE IF NOT EXISTS "landing_testimonial_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_testimonial_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_testimonial_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_testimonial_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_testimonial_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_testimonial_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing testimonial content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_testimonial_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing testimonial content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_testimonial_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_testimonial_content_exists
BEFORE UPDATE ON landing_package_items
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_testimonial_content_exists();

CREATE TRIGGER prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_testimonial_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_testimonial_content_reviews" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE,
    "reviewer" VARCHAR(100) NOT NULL,
    "age" INT NOT NULL,
    "address" VARCHAR(500) NOT NULL,
    "rating" "decimal_rating" NOT NULL,
    "review" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_testimonial_content_testimonials_id_pkey" PRIMARY KEY ("id")
);

-- migrate:down

DROP TABLE IF EXISTS "landing_testimonial_content_reviews";

DROP TRIGGER IF EXISTS "prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted" ON "landing_testimonial_content";
DROP TRIGGER IF EXISTS "prevent_delete_landing_section_header_if_landing_testimonial_content_exists" ON "landing_package_items";
DROP FUNCTION IF EXISTS "prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted"();
DROP FUNCTION IF EXISTS "prevent_delete_landing_section_header_if_landing_testimonial_content_exists"();

DROP TABLE IF EXISTS "landing_testimonial_content";

DROP DOMAIN IF EXISTS "decimal_rating";

