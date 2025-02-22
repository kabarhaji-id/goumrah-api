-- migrate:up

CREATE OR REPLACE FUNCTION set_airline_logo_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE airlines SET logo_id = NULL WHERE logo_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_airline_if_logo_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.logo_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.logo_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert airline with soft deleted logo'
                USING ERRCODE = '23503', CONSTRAINT = 'airlines_logo_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_airline_logo_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_airline_logo_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_airline_if_logo_is_soft_deleted
BEFORE INSERT OR UPDATE ON airlines
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_airline_if_logo_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_package_thumbnail_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE packages SET thumbnail_id = NULL WHERE thumbnail_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_if_thumbnail_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.thumbnail_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.thumbnail_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package with soft deleted thumbnail'
                USING ERRCODE = '23503', CONSTRAINT = 'packages_thumbnail_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_package_thumbnail_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_package_thumbnail_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_package_if_thumbnail_is_soft_deleted
BEFORE INSERT OR UPDATE ON packages
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_if_thumbnail_is_soft_deleted();

-- migrate:down

DROP TRIGGER prevent_insert_package_if_thumbnail_is_soft_deleted ON packages;
DROP TRIGGER set_package_thumbnail_id_null_on_image_soft_deleted ON images;
DROP FUNCTION prevent_insert_package_if_thumbnail_is_soft_deleted;
DROP FUNCTION set_package_thumbnail_id_null_on_image_soft_deleted;

DROP TRIGGER prevent_insert_airline_if_logo_is_soft_deleted ON airlines;
DROP TRIGGER set_airline_logo_id_null_on_image_soft_deleted ON images;
DROP FUNCTION prevent_insert_airline_if_logo_is_soft_deleted;
DROP FUNCTION set_airline_logo_id_null_on_image_soft_deleted;