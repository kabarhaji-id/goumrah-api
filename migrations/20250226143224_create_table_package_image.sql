-- migrate:up

CREATE TABLE IF NOT EXISTS "package_images" (
    "package_id" BIGINT NOT NULL,
    "image_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "package_images_id_pkey" PRIMARY KEY ("package_id", "image_id"),
    CONSTRAINT "package_images_package_id_fkey" FOREIGN KEY ("package_id") REFERENCES "packages" ("id"),
    CONSTRAINT "package_images_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE OR REPLACE FUNCTION delete_package_image_on_package_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_images SET deleted_at = NOW() WHERE package_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_image_on_package_soft_deleted
BEFORE UPDATE ON packages
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_image_on_package_soft_deleted();

CREATE OR REPLACE FUNCTION delete_package_image_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_image_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_image_on_image_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS delete_package_image_on_image_soft_deleted ON images;
DROP FUNCTION IF EXISTS delete_package_image_on_image_soft_deleted;

DROP TRIGGER IF EXISTS delete_package_image_on_package_soft_deleted ON packages;
DROP FUNCTION IF EXISTS delete_package_image_on_package_soft_deleted;

DROP TABLE IF EXISTS "package_images";
