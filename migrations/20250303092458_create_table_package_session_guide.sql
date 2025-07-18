-- migrate:up

CREATE TABLE IF NOT EXISTS "package_session_guides" (
    "package_session_id" BIGINT NOT NULL,
    "guide_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "package_session_guides_id_pkey" PRIMARY KEY ("package_session_id", "guide_id"),
    CONSTRAINT "package_session_guides_package_session_id_fkey" FOREIGN KEY ("package_session_id") REFERENCES "package_sessions" ("id"),
    CONSTRAINT "package_session_guides_guide_id_fkey" FOREIGN KEY ("guide_id") REFERENCES "guides" ("id")
);

CREATE OR REPLACE FUNCTION delete_package_session_guide_on_package_session_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_session_guides SET deleted_at = NOW() WHERE package_session_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_guide_if_package_session_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM package_sessions WHERE id = NEW.package_session_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session guide with soft deleted package session'
            USING ERRCODE = '23503', CONSTRAINT = 'package_session_guides_package_session_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_session_guide_on_package_session_soft_deleted
BEFORE UPDATE ON package_sessions
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_session_guide_on_package_session_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_guide_if_package_session_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_session_guides
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_guide_if_package_session_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_package_session_guide_on_guide_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_session_guides SET deleted_at = NOW() WHERE guide_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_guide_if_guide_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM guides WHERE id = NEW.guide_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session guide with soft deleted guide'
            USING ERRCODE = '23503', CONSTRAINT = 'package_session_guides_guide_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_session_guide_on_guide_soft_deleted
BEFORE UPDATE ON guides
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_session_guide_on_guide_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_guide_if_guide_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_session_guides
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_guide_if_guide_is_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS prevent_insert_package_session_guide_if_guide_is_soft_deleted ON package_session_guides;
DROP FUNCTION IF EXISTS prevent_insert_package_session_guide_if_guide_is_soft_deleted;
DROP TRIGGER IF EXISTS delete_package_session_guide_on_guide_soft_deleted ON guides;
DROP FUNCTION IF EXISTS delete_package_session_guide_on_guide_soft_deleted;

DROP TRIGGER IF EXISTS prevent_insert_package_session_guide_if_package_session_is_soft_deleted ON package_session_guides;
DROP FUNCTION IF EXISTS prevent_insert_package_session_guide_if_package_session_is_soft_deleted;
DROP TRIGGER IF EXISTS delete_package_session_guide_on_package_session_soft_deleted ON package_sessions;
DROP FUNCTION IF EXISTS delete_package_session_guide_on_package_session_soft_deleted;

DROP TABLE IF EXISTS "package_session_guides";
