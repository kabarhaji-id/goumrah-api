-- migrate:up

CREATE TYPE "guide_type" AS ENUM ('Perjalanan', 'Ibadah');

CREATE TABLE IF NOT EXISTS "guides" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "avatar_id" BIGINT NULL DEFAULT NULL,
    "name" VARCHAR(100) NOT NULL,
    "type" "guide_type" NOT NULL,
    "description" VARCHAR(500) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "guides_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "guides_avatar_id_fkey" FOREIGN KEY ("avatar_id") REFERENCES "images" ("id")
);

CREATE UNIQUE INDEX "guides_name_unique"
ON "guides" (UPPER("name"))
WHERE "deleted_at" IS NULL;

CREATE OR REPLACE FUNCTION set_guide_avatar_id_null_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE guides SET avatar_id = NULL WHERE avatar_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_guide_if_avatar_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.avatar_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.avatar_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert guide with soft deleted avatar'
                USING ERRCODE = '23503', CONSTRAINT = 'guides_avatar_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_guide_avatar_id_null_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_guide_avatar_id_null_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_guide_if_avatar_is_soft_deleted
BEFORE INSERT OR UPDATE ON guides
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_guide_if_avatar_is_soft_deleted();

-- migrate:down

DROP TRIGGER prevent_insert_guide_if_avatar_is_soft_deleted ON guides;
DROP TRIGGER set_guide_avatar_id_null_on_image_soft_deleted ON images;
DROP FUNCTION prevent_insert_guide_if_avatar_is_soft_deleted;
DROP FUNCTION set_guide_avatar_id_null_on_image_soft_deleted;

DROP TABLE IF EXISTS "guides";

DROP TYPE IF EXISTS "guide_type";