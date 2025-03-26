-- migrate:up

ALTER TABLE IF EXISTS "package_sessions"
ADD COLUMN "double_price" DECIMAL(10, 2) NOT NULL DEFAULT 0,
ADD COLUMN "double_final_price" DECIMAL(10, 2) NULL DEFAULT NULL,
ADD COLUMN "triple_price" DECIMAL(10, 2) NOT NULL DEFAULT 0,
ADD COLUMN "triple_final_price" DECIMAL(10, 2) NULL DEFAULT NULL,
ADD COLUMN "quad_price" DECIMAL(10, 2) NOT NULL DEFAULT 0,
ADD COLUMN "quad_final_price" DECIMAL(10, 2) NULL DEFAULT NULL,
ADD COLUMN "infant_price" DECIMAL(10, 2) NULL DEFAULT NULL,
ADD COLUMN "infant_final_price" DECIMAL(10, 2) NULL DEFAULT NULL,
ADD COLUMN "quota" INT NOT NULL DEFAULT 0;

-- migrate:down

ALTER TABLE IF EXISTS "package_sessions"
DROP COLUMN "double_price",
DROP COLUMN "double_final_price",
DROP COLUMN "triple_price",
DROP COLUMN "triple_final_price",
DROP COLUMN "quad_price",
DROP COLUMN "quad_final_price",
DROP COLUMN "infant_price",
DROP COLUMN "infant_final_price",
DROP COLUMN "quota";
