-- migrate:up

ALTER TABLE IF EXISTS "landing_hero_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_section_headers"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_package_items"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_single_package_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_package_details"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_package_detail_items"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_packages_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_features_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_features_content_benefits"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_moments_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_moments_content_images"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_affiliates_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_affiliates_content_affiliates"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_faq_content"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_faq_content_faqs"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE IF EXISTS "landing_menus"
ADD COLUMN "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
ADD COLUMN "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE;

-- migrate:down

ALTER TABLE IF EXISTS "landing_hero_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_section_headers"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_package_items"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_single_package_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_package_details"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_package_detail_items"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_packages_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_features_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_features_content_benefits"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_moments_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_moments_content_images"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_affiliates_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_affiliates_content_affiliates"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_faq_content"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_faq_content_faqs"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";

ALTER TABLE IF EXISTS "landing_menus"
DROP COLUMN "is_mobile",
DROP COLUMN "is_desktop";
