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

CREATE TRIGGER delete_addon_on_category_soft_deleted
BEFORE UPDATE ON addon_categories
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_addon_on_category_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS delete_addon_on_category_soft_deleted ON addon_categories;
DROP FUNCTION IF EXISTS delete_addon_on_category_soft_deleted;

DROP TABLE IF EXISTS "addons";