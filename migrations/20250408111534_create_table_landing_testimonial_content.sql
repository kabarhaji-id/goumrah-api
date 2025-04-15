-- migrate:up

CREATE DOMAIN "decimal_rating" AS DECIMAL NOT NULL CHECK (value BETWEEN 1 AND 5);

CREATE TABLE IF NOT EXISTS "landing_testimonial_content" (
    "id" INT NOT NULL DEFAULT (1),
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE,
    "landing_section_header_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_testimonial_content_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "landing_testimonial_content_id_check" CHECK ("id" = 1),
    CONSTRAINT "landing_testimonial_content_landing_section_header_id_fkey" FOREIGN KEY ("landing_section_header_id") REFERENCES "landing_section_headers" ("id")
);

CREATE OR REPLACE FUNCTION prevent_delete_landing_section_header_if_landing_testimonial_content_exists()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_testimonial_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing testimonial content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_testimonial_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing testimonial content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_testimonial_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_testimonial_content_exists
BEFORE UPDATE ON landing_section_headers
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION prevent_delete_landing_section_header_if_landing_testimonial_content_exists();

CREATE TRIGGER prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted
BEFORE INSERT OR UPDATE ON landing_testimonial_content
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted();

CREATE TABLE IF NOT EXISTS "landing_testimonial_content_reviews" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "is_enabled" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_mobile" BOOLEAN NOT NULL DEFAULT TRUE,
    "is_desktop" BOOLEAN NOT NULL DEFAULT TRUE,
    "reviewer" VARCHAR(100) NOT NULL,
    "age" INT NOT NULL,
    "address" VARCHAR(500) NOT NULL,
    "rating" "decimal_rating" NOT NULL,
    "review" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "landing_testimonial_content_testimonials_id_pkey" PRIMARY KEY ("id")
);

INSERT INTO "landing_section_headers" (
    "title", "subtitle", "tags_line"
) VALUES
    ('Testimoni jama''ah', 'Apa yang para Jama''ah katakan tentang kami', NULL);

INSERT INTO "landing_testimonial_content" (
    "id", "is_enabled", "is_mobile", "is_desktop", "landing_section_header_id"
) VALUES
    (1, TRUE, TRUE, TRUE, 9);

INSERT INTO "landing_testimonial_content_reviews" (
    "is_enabled", "is_mobile", "is_desktop", "reviewer", "age", "address", "rating", "review"
) VALUES
    (TRUE, TRUE, TRUE, 'Ahmad Ali', 45, 'Jakarta, Indonesia', 5.0, 'Paket umrah ini benar-benar luar biasa. Akomodasi sangat nyaman, dan pemandu sangat berpengalaman serta membantu. Sangat direkomendasikan!'),
    (TRUE, TRUE, TRUE, 'Fatimah Noor', 38, 'Bandung, Indonesia', 4.0, 'Secara keseluruhan pengalaman yang sangat baik. Transportasi lancar dan staf sangat ramah. Namun, pilihan makanannya bisa lebih ditingkatkan.'),
    (TRUE, TRUE, TRUE, 'Yusuf Rahman', 50, 'Surabaya, Indonesia', 5.0, 'Perjalanan spiritual yang benar-benar luar biasa dan bebas stres berkat paket ini. Semua diatur dengan sempurna dari awal hingga akhir.'),
    (TRUE, TRUE, TRUE, 'Aisyah Malik', 30, 'Medan, Indonesia', 4.5, 'Pelayanan luar biasa! Pemimpin grup sangat membantu dan memastikan semua jamaah memiliki pengalaman yang lancar. Sangat saya rekomendasikan.'),
    (TRUE, TRUE, TRUE, 'Muhammad Hasan', 60, 'Makassar, Indonesia', 5.0, 'Layanan yang sangat baik! Mulai dari pengurusan visa hingga akomodasi, semuanya ditangani dengan sangat profesional.'),
    (TRUE, TRUE, TRUE, 'Siti Amina', 42, 'Yogyakarta, Indonesia', 4.5, 'Paket ini sangat cocok untuk keluarga. Fasilitasnya bersih, dan pemandu sangat ramah. Jadwalnya juga tidak terlalu padat, sehingga cukup fleksibel.'),
    (TRUE, TRUE, TRUE, 'Ali Fauzan', 55, 'Balikpapan, Indonesia', 5.0, 'Saya sangat puas dengan layanan ini. Semua kebutuhan saya selama umrah terpenuhi dengan baik. Terima kasih telah membuat perjalanan ini begitu berkesan.'),
    (TRUE, TRUE, TRUE, 'Nurul Hidayah', 36, 'Palembang, Indonesia', 4.7, 'Layanan yang sangat profesional. Timnya sangat responsif dan cepat membantu jika ada kebutuhan mendadak. Sungguh pengalaman yang tak terlupakan.'),
    (TRUE, TRUE, TRUE, 'Rizki Aditya', 29, 'Denpasar, Indonesia', 4.8, 'Sebagai jamaah muda, saya merasa sangat terbantu dengan panduan dan fleksibilitas jadwal yang diberikan. Pengalaman spiritual yang sangat luar biasa.'),
    (TRUE, TRUE, TRUE, 'Zahra Putri', 47, 'Banda Aceh, Indonesia', 5.0, 'Sangat puas! Dari mulai keberangkatan hingga pulang, semuanya diurus dengan sangat rapi. Hotelnya nyaman, dan lokasi strategis dekat Haram.');

-- migrate:down

DROP TABLE IF EXISTS "landing_testimonial_content_reviews";

DROP TRIGGER IF EXISTS "prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted" ON "landing_testimonial_content";
DROP TRIGGER IF EXISTS "prevent_delete_landing_section_header_if_landing_testimonial_content_exists" ON "landing_section_headers";
DROP FUNCTION IF EXISTS "prevent_insert_landing_testimonial_content_if_landing_section_header_is_soft_deleted"();
DROP FUNCTION IF EXISTS "prevent_delete_landing_section_header_if_landing_testimonial_content_exists"();

DROP TABLE IF EXISTS "landing_testimonial_content";

DROP DOMAIN IF EXISTS "decimal_rating";

