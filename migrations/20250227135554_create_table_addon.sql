-- migrate:up

CREATE TABLE IF NOT EXISTS "addons" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "category_id" BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "price" DECIMAL(13, 2) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT "addons_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "addons_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "addon_categories" ("id")
);

CREATE UNIQUE INDEX "addons_name_unique"
ON "addons" (UPPER("name"))
WHERE "deleted_at" IS NULL;

CREATE OR REPLACE FUNCTION delete_addon_on_category_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE addons SET deleted_at = NOW() WHERE category_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_addon_if_category_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM addon_categories WHERE id = NEW.category_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert addon with soft deleted category'
            USING ERRCODE = '23503', CONSTRAINT = 'addons_category_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_addon_on_category_soft_deleted
BEFORE UPDATE ON addon_categories
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_addon_on_category_soft_deleted();

CREATE TRIGGER prevent_insert_addon_if_category_is_soft_deleted
BEFORE INSERT OR UPDATE ON addons
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_addon_if_category_is_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS prevent_insert_addon_if_category_is_soft_deleted ON addons;
DROP TRIGGER IF EXISTS delete_addon_on_category_soft_deleted ON addon_categories;
DROP FUNCTION IF EXISTS prevent_insert_addon_if_category_is_soft_deleted;
DROP FUNCTION IF EXISTS delete_addon_on_category_soft_deleted;

DROP TABLE IF EXISTS "addons";