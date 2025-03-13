-- migrate:up

ALTER TABLE IF EXISTS "package_sessions"
ADD COLUMN "bus_id" BIGINT NULL DEFAULT NULL,
ADD CONSTRAINT "package_sessions_bus_id_fkey" FOREIGN KEY ("bus_id") REFERENCES "buses" ("id");

CREATE OR REPLACE FUNCTION delete_package_session_on_bus_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET deleted_at = NOW() WHERE bus_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_if_bus_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM buses WHERE id = NEW.bus_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session with soft deleted bus'
            USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_bus_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_session_on_bus_soft_deleted
BEFORE UPDATE ON buses
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_session_on_bus_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_if_bus_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_sessions
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_if_bus_is_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS "prevent_insert_package_session_if_bus_is_soft_deleted" ON "package_sessions";
DROP FUNCTION IF EXISTS "prevent_insert_package_session_if_bus_is_soft_deleted";
DROP TRIGGER IF EXISTS "delete_package_session_on_bus_soft_deleted" ON "buses";
DROP FUNCTION IF EXISTS "delete_package_session_on_bus_soft_deleted";

ALTER TABLE IF EXISTS "package_sessions"
DROP COLUMN IF EXISTS "bus_id",
DROP CONSTRAINT IF EXISTS "package_sessions_bus_id_fkey";
