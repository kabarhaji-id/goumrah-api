-- migrate:up

CREATE TABLE IF NOT EXISTS "itinerary_widget_activities" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "description" VARCHAR(500) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_activities_id_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widget_activity_images" (
    "itinerary_widget_activity_id" BIGINT NOT NULL,
    "image_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_activity_images_itinerary_widget_activity_id_image_id_key" PRIMARY KEY ("itinerary_widget_activity_id", "image_id"),
    CONSTRAINT "itinerary_widget_activity_images_itinerary_widget_activity_id_fkey" FOREIGN KEY ("itinerary_widget_activity_id") REFERENCES "itinerary_widget_activities" ("id"),
    CONSTRAINT "itinerary_widget_activity_images_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widget_hotels" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "hotel_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_hotels_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "itinerary_widget_hotels_hotel_id_fkey" FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widget_informations" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "description" VARCHAR(500) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_informations_id_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widget_transports" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "transportation" VARCHAR(100) NOT NULL,
    "from" VARCHAR(100) NOT NULL,
    "to" VARCHAR(100) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_transports_id_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widget_recommendations" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "description" VARCHAR(500) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_recommendations_id_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widget_recommendation_images" (
    "itinerary_widget_recommendation_id" BIGINT NOT NULL,
    "image_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widget_recommendation_images_pkey" PRIMARY KEY ("itinerary_widget_recommendation_id", "image_id"),
    CONSTRAINT "itinerary_widget_recommendation_images_fkey" FOREIGN KEY ("itinerary_widget_recommendation_id") REFERENCES "itinerary_widget_recommendations" ("id"),
    CONSTRAINT "itinerary_widget_recommendation_images_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_widgets" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "activity_id" BIGINT NULL DEFAULT NULL,
    "hotel_id" BIGINT NULL DEFAULT NULL,
    "information_id" BIGINT NULL DEFAULT NULL,
    "transport_id" BIGINT NULL DEFAULT NULL,
    "recommendation_id" BIGINT NULL DEFAULT NULL,
    "next_id" BIGINT NULL DEFAULT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_widgets_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "itinerary_widgets_activity_id_fkey" FOREIGN KEY ("activity_id") REFERENCES "itinerary_widget_activities" ("id"),
    CONSTRAINT "itinerary_widgets_hotel_id_fkey" FOREIGN KEY ("hotel_id") REFERENCES "itinerary_widget_hotels" ("id"),
    CONSTRAINT "itinerary_widgets_information_id_fkey" FOREIGN KEY ("information_id") REFERENCES "itinerary_widget_informations" ("id"),
    CONSTRAINT "itinerary_widgets_transport_id_fkey" FOREIGN KEY ("transport_id") REFERENCES "itinerary_widget_transports" ("id"),
    CONSTRAINT "itinerary_widgets_recommendation_id_fkey" FOREIGN KEY ("recommendation_id") REFERENCES "itinerary_widget_recommendations" ("id"),
    CONSTRAINT "itinerary_widgets_next_id_fkey" FOREIGN KEY ("next_id") REFERENCES "itinerary_widgets" ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_days" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "description" VARCHAR(500) NOT NULL,
    "widget_id" BIGINT NULL DEFAULT NULL,
    "next_id" BIGINT NULL DEFAULT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_days_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "itinerary_days_widget_id_fkey" FOREIGN KEY ("widget_id") REFERENCES "itinerary_widgets" ("id"),
    CONSTRAINT "itinerary_days_next_id_fkey" FOREIGN KEY ("next_id") REFERENCES "itinerary_days" ("id")
);

CREATE TABLE IF NOT EXISTS "itineraries" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    "city" VARCHAR(100) NOT NULL,
    "day_id" BIGINT NOT NULL,
    "next_id" BIGINT NULL DEFAULT NULL,
    
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itineraries_id_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "itineraries_day_id_fkey" FOREIGN KEY ("day_id") REFERENCES "itinerary_days" ("id"),
    CONSTRAINT "itineraries_next_id_fkey" FOREIGN KEY ("next_id") REFERENCES "itineraries" ("id")
);

CREATE TABLE IF NOT EXISTS "itinerary_images" (
    "itinerary_id" BIGINT NOT NULL,
    "image_id" BIGINT NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT "itinerary_images_itinerary_id_image_id_key" PRIMARY KEY ("itinerary_id", "image_id"),
    CONSTRAINT "itinerary_images_itinerary_id_fkey" FOREIGN KEY ("itinerary_id") REFERENCES "itineraries" ("id"),
    CONSTRAINT "itinerary_images_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "images" ("id")
);

ALTER TABLE IF EXISTS "package_sessions"
ADD COLUMN "itinerary_id" BIGINT NULL DEFAULT NULL,
ADD CONSTRAINT "package_sessions_itinerary_id_fkey" FOREIGN KEY ("itinerary_id") REFERENCES "itineraries" ("id");

-- Itinerary Widget Activity Image

CREATE OR REPLACE FUNCTION delete_itinerary_widget_activity_image_on_itinerary_widget_activity_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_activity_images SET deleted_at = NOW() WHERE itinerary_widget_activity_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_activity_image_if_itinerary_widget_activity_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM itineraries WHERE id = NEW.itinerary_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget activity image with soft deleted itinerary widget activity'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_activity_images_itinerary_widget_activity_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_widget_activity_image_on_itinerary_widget_activity_soft_deleted
BEFORE UPDATE ON itinerary_widget_activities
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_widget_activity_image_on_itinerary_widget_activity_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_activity_image_if_itinerary_widget_activity_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widget_activity_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_activity_image_if_itinerary_widget_activity_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_itinerary_widget_activity_image_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_activity_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_activity_image_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget activity image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_activity_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_widget_activity_image_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_widget_activity_image_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_activity_image_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widget_activity_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_activity_image_if_image_is_soft_deleted();

-- Itinerary Widget Hotel

CREATE OR REPLACE FUNCTION delete_itinerary_widget_hotel_on_hotel_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_hotels SET deleted_at = NOW() WHERE hotel_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM hotels WHERE id = NEW.hotel_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget hotel with soft deleted hotel'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_hotels_hotel_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_widget_hotel_on_hotel_soft_deleted
BEFORE UPDATE ON hotels
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_widget_hotel_on_hotel_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widget_hotels
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted();

-- Itinerary Widget Recommendation Image

CREATE OR REPLACE FUNCTION delete_itinerary_widget_recommendation_image_on_itinerary_widget_recommendation_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_recommendation_images SET deleted_at = NOW() WHERE itinerary_widget_recommendation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_recommendation_image_if_itinerary_widget_recommendation_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM itineraries WHERE id = NEW.itinerary_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget recommendation image with soft deleted itinerary widget recommendation'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_recommendation_images_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_widget_recommendation_image_on_itinerary_widget_recommendation_soft_deleted
BEFORE UPDATE ON itinerary_widget_recommendations
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_widget_recommendation_image_on_itinerary_widget_recommendation_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_recommendation_image_if_itinerary_widget_recommendation_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widget_recommendation_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_recommendation_image_if_itinerary_widget_recommendation_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_itinerary_widget_recommendation_image_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_recommendation_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_recommendation_image_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget recommendation image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_recommendation_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_widget_recommendation_image_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_widget_recommendation_image_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_recommendation_image_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widget_recommendation_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_recommendation_image_if_image_is_soft_deleted();

-- Itinerary Widget

CREATE OR REPLACE FUNCTION set_activity_id_null_on_activity_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET activity_id = NULL WHERE activity_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_if_activity_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.activity_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_activities WHERE id = NEW.activity_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted activity'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_activity_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_activity_id_null_on_activity_soft_deleted
BEFORE UPDATE ON itinerary_widget_activities
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_activity_id_null_on_activity_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_if_activity_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widgets
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_if_activity_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_hotel_id_null_on_hotel_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET hotel_id = NULL WHERE hotel_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_if_hotel_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.hotel_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_hotels WHERE id = NEW.hotel_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted hotel'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_hotel_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_hotel_id_null_on_hotel_soft_deleted
BEFORE UPDATE ON itinerary_widget_hotels
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_hotel_id_null_on_hotel_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_if_hotel_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widgets
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_if_hotel_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_information_id_null_on_information_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET information_id = NULL WHERE information_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_if_information_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.information_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_informations WHERE id = NEW.information_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted information'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_information_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_information_id_null_on_information_soft_deleted
BEFORE UPDATE ON itinerary_widget_informations
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_information_id_null_on_information_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_if_information_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widgets
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_if_information_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_transport_id_null_on_transport_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET transport_id = NULL WHERE transport_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_if_transport_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.transport_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_transports WHERE id = NEW.transport_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted transport'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_transport_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_transport_id_null_on_transport_soft_deleted
BEFORE UPDATE ON itinerary_widget_transports
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_transport_id_null_on_transport_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_if_transport_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widgets
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_if_transport_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_recommendation_id_null_on_recommendation_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET recommendation_id = NULL WHERE recommendation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_if_recommendation_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.recommendation_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_recommendations WHERE id = NEW.recommendation_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted recommendation'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_recommendation_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_recommendation_id_null_on_recommendation_soft_deleted
BEFORE UPDATE ON itinerary_widget_recommendations
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_recommendation_id_null_on_recommendation_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_if_recommendation_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widgets
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_if_recommendation_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_next_id_null_on_itinerary_widget_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_widget_if_next_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widgets WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_next_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_next_id_null_on_itinerary_widget_soft_deleted
BEFORE UPDATE ON itinerary_widgets
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_next_id_null_on_itinerary_widget_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_widget_if_next_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_widgets
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_widget_if_next_is_soft_deleted();

-- Itinerary Day

CREATE OR REPLACE FUNCTION set_widget_id_null_on_widget_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_days SET widget_id = NULL WHERE widget_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_day_if_widget_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.widget_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widgets WHERE id = NEW.widget_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary day with soft deleted widget'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_days_widget_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_widget_id_null_on_widget_soft_deleted
BEFORE UPDATE ON itinerary_widgets
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_widget_id_null_on_widget_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_day_if_widget_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_days
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_day_if_widget_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_next_id_null_on_itinerary_day_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_days SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_day_if_next_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_days WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary day with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_days_next_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_next_id_null_on_itinerary_day_soft_deleted
BEFORE UPDATE ON itinerary_days
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_next_id_null_on_itinerary_day_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_day_if_next_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_days
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_day_if_next_is_soft_deleted();

-- Itinerary

CREATE OR REPLACE FUNCTION delete_itinerary_on_day_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itineraries SET deleted_at = NOW() WHERE day_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_if_day_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM itinerary_days WHERE id = NEW.day_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary with soft deleted day'
            USING ERRCODE = '23503', CONSTRAINT = 'itineraries_day_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_on_day_soft_deleted
BEFORE UPDATE ON itinerary_days
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_on_day_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_if_day_is_soft_deleted
BEFORE INSERT OR UPDATE ON itineraries
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_if_day_is_soft_deleted();

CREATE OR REPLACE FUNCTION set_next_id_null_on_itinerary_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itineraries SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_if_next_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itineraries WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'itineraries_next_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_next_id_null_on_itinerary_soft_deleted
BEFORE UPDATE ON itineraries
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION set_next_id_null_on_itinerary_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_if_next_is_soft_deleted
BEFORE INSERT OR UPDATE ON itineraries
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_if_next_is_soft_deleted();

-- Itinerary Image

CREATE OR REPLACE FUNCTION delete_itinerary_image_on_itinerary_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_images SET deleted_at = NOW() WHERE itinerary_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_image_if_package_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM itineraries WHERE id = NEW.itinerary_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary image with soft deleted itinerary'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_images_itinerary_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_image_on_itinerary_soft_deleted
BEFORE UPDATE ON itineraries
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_image_on_itinerary_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_image_if_package_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_image_if_package_is_soft_deleted();

CREATE OR REPLACE FUNCTION delete_itinerary_image_on_image_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_insert_itinerary_image_if_image_is_soft_deleted()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_itinerary_image_on_image_soft_deleted
BEFORE UPDATE ON images
FOR EACH ROW
WHEN (OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL)
EXECUTE FUNCTION delete_itinerary_image_on_image_soft_deleted();

CREATE TRIGGER prevent_insert_itinerary_image_if_image_is_soft_deleted
BEFORE INSERT OR UPDATE ON itinerary_images
FOR EACH ROW
EXECUTE FUNCTION prevent_insert_itinerary_image_if_image_is_soft_deleted();

-- migrate:down

DROP TRIGGER IF EXISTS prevent_insert_itinerary_image_if_image_is_soft_deleted ON itinerary_images;
DROP TRIGGER IF EXISTS delete_itinerary_image_on_image_soft_deleted ON images;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_image_if_image_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_image_on_image_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_image_if_package_is_soft_deleted ON itinerary_images;
DROP TRIGGER IF EXISTS delete_itinerary_image_on_itinerary_soft_deleted ON itineraries;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_image_if_package_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_image_on_itinerary_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_day_if_next_is_soft_deleted ON itinerary_days;
DROP TRIGGER IF EXISTS set_next_id_null_on_itinerary_day_soft_deleted ON itinerary_days;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_day_if_next_is_soft_deleted();
DROP FUNCTION IF EXISTS set_next_id_null_on_itinerary_day_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_if_next_is_soft_deleted ON itinerary_widgets;
DROP TRIGGER IF EXISTS set_next_id_null_on_itinerary_widget_soft_deleted ON itinerary_widgets;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_if_next_is_soft_deleted();
DROP FUNCTION IF EXISTS set_next_id_null_on_itinerary_widget_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_if_recommendation_is_soft_deleted ON itinerary_widgets;
DROP TRIGGER IF EXISTS set_recommendation_id_null_on_recommendation_soft_deleted ON itinerary_widget_recommendations;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_if_recommendation_is_soft_deleted();
DROP FUNCTION IF EXISTS set_recommendation_id_null_on_recommendation_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_if_transport_is_soft_deleted ON itinerary_widgets;
DROP TRIGGER IF EXISTS set_transport_id_null_on_transport_soft_deleted ON itinerary_widget_transports;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_if_transport_is_soft_deleted();
DROP FUNCTION IF EXISTS set_transport_id_null_on_transport_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_if_information_is_soft_deleted ON itinerary_widgets;
DROP TRIGGER IF EXISTS set_information_id_null_on_information_soft_deleted ON itinerary_widget_informations;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_if_information_is_soft_deleted();
DROP FUNCTION IF EXISTS set_information_id_null_on_information_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_if_hotel_is_soft_deleted ON itinerary_widgets;
DROP TRIGGER IF EXISTS delete_itinerary_widget_hotel_on_hotel_soft_deleted ON hotels;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_if_hotel_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_widget_hotel_on_hotel_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_if_activity_is_soft_deleted ON itinerary_widgets;
DROP TRIGGER IF EXISTS set_activity_id_null_on_activity_soft_deleted ON itinerary_widget_activities;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_if_activity_is_soft_deleted();
DROP FUNCTION IF EXISTS set_activity_id_null_on_activity_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_activity_image_if_image_is_soft_deleted ON itinerary_widget_activity_images;
DROP TRIGGER IF EXISTS delete_itinerary_widget_activity_image_on_image_soft_deleted ON images;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_activity_image_if_image_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_widget_activity_image_on_image_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_activity_image_if_itinerary_widget_activity_is_soft_deleted ON itinerary_widget_activity_images;
DROP TRIGGER IF EXISTS delete_itinerary_widget_activity_image_on_itinerary_widget_activity_soft_deleted ON itinerary_widget_activities;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_activity_image_if_itinerary_widget_activity_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_widget_activity_image_on_itinerary_widget_activity_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_recommendation_image_if_image_is_soft_deleted ON itinerary_widget_recommendation_images;
DROP TRIGGER IF EXISTS delete_itinerary_widget_recommendation_image_on_image_soft_deleted ON images;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_recommendation_image_if_image_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_widget_recommendation_image_on_image_soft_deleted();

DROP TRIGGER IF EXISTS prevent_insert_itinerary_widget_recommendation_image_if_itinerary_widget_recommendation_is_soft_deleted ON itinerary_widget_recommendation_images;
DROP TRIGGER IF EXISTS delete_itinerary_widget_recommendation_image_on_itinerary_widget_recommendation_soft_deleted ON itinerary_widget_recommendations;
DROP FUNCTION IF EXISTS prevent_insert_itinerary_widget_recommendation_image_if_itinerary_widget_recommendation_is_soft_deleted();
DROP FUNCTION IF EXISTS delete_itinerary_widget_recommendation_image_on_itinerary_widget_recommendation_soft_deleted();

ALTER TABLE IF EXISTS "package_sessions"
DROP COLUMN IF EXISTS "itinerary_id",
DROP CONSTRAINT IF EXISTS "package_sessions_itinerary_id_fkey";

DROP TABLE IF EXISTS itinerary_images;

DROP TABLE IF EXISTS itineraries;

DROP TABLE IF EXISTS itinerary_days;

DROP TABLE IF EXISTS itinerary_widgets;

DROP TABLE IF EXISTS itinerary_widget_recommendation_images;

DROP TABLE IF EXISTS itinerary_widget_recommendations;

DROP TABLE IF EXISTS itinerary_widget_transports;

DROP TABLE IF EXISTS itinerary_widget_informations;

DROP TABLE IF EXISTS itinerary_widget_hotels;

DROP TABLE IF EXISTS itinerary_widget_activity_images;

DROP TABLE IF EXISTS itinerary_widget_activities;
