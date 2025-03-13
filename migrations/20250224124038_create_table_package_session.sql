-- migrate:up

CREATE TABLE IF NOT EXISTS "package_sessions" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "package_id" BIGINT NOT NULL,
    "embarkation_id" BIGINT NOT NULL,
    "departure_date" DATE NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "package_sessions_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "package_sessions_package_id_fkey" FOREIGN KEY ("package_id") REFERENCES "packages" ("id"),
    CONSTRAINT "package_sessions_embarkation_id_fkey" FOREIGN KEY ("embarkation_id") REFERENCES "embarkations" ("id")
);

CREATE UNIQUE INDEX "package_sessions_package_id_embarkation_id_departure_date_unique"
ON "package_sessions" ("package_id", "embarkation_id", "departure_date")
WHERE "deleted_at" IS NULL;

CREATE OR REPLACE FUNCTION delete_package_session_on_embarkation_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET deleted_at = NOW() WHERE embarkation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_if_embarkation_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM embarkations WHERE id = NEW.embarkation_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session with soft deleted embarkation'
            USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_embarkation_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_session_on_embarkation_soft_deleted
BEFORE UPDATE ON embarkations
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_session_on_embarkation_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_if_embarkation_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_sessions
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_if_embarkation_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_package_session_on_package_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET deleted_at = NOW() WHERE embarkation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_package_session_if_package_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.package_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM packages WHERE id = NEW.package_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package session with soft deleted package'
                USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_package_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_package_session_on_package_soft_deleted
BEFORE UPDATE ON packages
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_package_session_on_package_soft_deleted();

CREATE TRIGGER prevent_insert_package_session_if_package_is_soft_deleted
BEFORE INSERT OR UPDATE ON package_sessions
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_package_session_if_package_is_soft_deleted();

-- migrate:down

DROP TRIGGER prevent_insert_package_session_if_package_is_soft_deleted ON package_sessions;
DROP TRIGGER delete_package_session_on_package_soft_deleted ON packages;
DROP FUNCTION prevent_insert_package_session_if_package_is_soft_deleted;
DROP FUNCTION delete_package_session_on_package_soft_deleted;

DROP TRIGGER prevent_insert_package_session_if_embarkation_is_soft_deleted ON package_sessions;
DROP TRIGGER delete_package_session_on_embarkation_soft_deleted ON embarkations;
DROP FUNCTION prevent_insert_package_session_if_embarkation_is_soft_deleted;
DROP FUNCTION delete_package_session_on_embarkation_soft_deleted;

DROP TABLE IF EXISTS "package_sessions";

DROP TYPE IF EXISTS "package_session_type";