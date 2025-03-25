SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: bus_class; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.bus_class AS ENUM (
    'Economy',
    'VIP'
);


--
-- Name: flight_class; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.flight_class AS ENUM (
    'Economy',
    'Business',
    'First'
);


--
-- Name: guide_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.guide_type AS ENUM (
    'Perjalanan',
    'Ibadah'
);


--
-- Name: package_category; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.package_category AS ENUM (
    'Silver',
    'Gold',
    'Platinum',
    'Luxury'
);


--
-- Name: package_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.package_type AS ENUM (
    'Reguler',
    'Plus'
);


--
-- Name: rating; Type: DOMAIN; Schema: public; Owner: -
--

CREATE DOMAIN public.rating AS smallint NOT NULL
	CONSTRAINT rating_check CHECK (((VALUE >= 1) AND (VALUE <= 5)));


--
-- Name: skytrax_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.skytrax_type AS ENUM (
    'Full Service',
    'Low Cost'
);


--
-- Name: user_role; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.user_role AS ENUM (
    'REGISTERED_USER',
    'CUSTOMER',
    'TRAVEL_AGENT',
    'ADMINISTRATOR',
    'CUSTOMER_SUPPORT',
    'PARTNER'
);


--
-- Name: delete_addon_on_category_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_addon_on_category_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE addons SET deleted_at = NOW() WHERE category_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_flight_on_airline_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_flight_on_airline_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flights SET deleted_at = NOW() WHERE airline_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_flight_on_arrival_airport_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_flight_on_arrival_airport_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flights SET deleted_at = NOW() WHERE arrival_airport_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_flight_on_departure_airport_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_flight_on_departure_airport_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flights SET deleted_at = NOW() WHERE departure_airport_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_flight_route_on_flight_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_flight_route_on_flight_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE flight_routes SET deleted_at = NOW() WHERE flight_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_hotel_image_on_hotel_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_hotel_image_on_hotel_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE hotel_images SET deleted_at = NOW() WHERE hotel_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_hotel_image_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_hotel_image_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE hotel_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_image_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_image_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_image_on_itinerary_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_image_on_itinerary_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_images SET deleted_at = NOW() WHERE itinerary_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_on_day_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_on_day_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itineraries SET deleted_at = NOW() WHERE day_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_widget_activity_image_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_widget_activity_image_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_activity_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_widget_activity_image_on_itinerary_widget_acti(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_widget_activity_image_on_itinerary_widget_acti() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_activity_images SET deleted_at = NOW() WHERE itinerary_widget_activity_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_widget_hotel_on_hotel_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_widget_hotel_on_hotel_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_hotels SET deleted_at = NOW() WHERE hotel_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_widget_recommendation_image_on_image_soft_dele(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_widget_recommendation_image_on_image_soft_dele() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_recommendation_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_itinerary_widget_recommendation_image_on_itinerary_widge(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_itinerary_widget_recommendation_image_on_itinerary_widge() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widget_recommendation_images SET deleted_at = NOW() WHERE itinerary_widget_recommendation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_landing_moments_content_image_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_landing_moments_content_image_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        DELETE FROM landing_moments_content_images WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_landing_package_detail_item_on_landing_package_detail_so(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_landing_package_detail_item_on_landing_package_detail_so() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_package_detail_items SET deleted_at = NOW() WHERE landing_package_detail_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_landing_package_detail_item_on_landing_package_item_soft(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_landing_package_detail_item_on_landing_package_item_soft() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_package_detail_items SET deleted_at = NOW() WHERE landing_package_item_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_landing_package_item_on_package_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_landing_package_item_on_package_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_package_items SET deleted_at = NOW() WHERE package_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_image_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_image_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_images SET deleted_at = NOW() WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_image_on_package_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_image_on_package_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_images SET deleted_at = NOW() WHERE package_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_session_guide_on_guide_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_session_guide_on_guide_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_session_guides SET deleted_at = NOW() WHERE guide_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_session_guide_on_package_session_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_session_guide_on_package_session_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_session_guides SET deleted_at = NOW() WHERE package_session_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_session_on_bus_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_session_on_bus_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET deleted_at = NOW() WHERE bus_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_session_on_embarkation_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_session_on_embarkation_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET deleted_at = NOW() WHERE embarkation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: delete_package_session_on_package_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.delete_package_session_on_package_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET deleted_at = NOW() WHERE embarkation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_delete_landing_package_detail_if_landing_packages_conte(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_package_detail_if_landing_packages_conte() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_packages_content WHERE silver_landing_package_detail_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing package detail with existing landing packages content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_silver_landing_package_detail_id_fkey';
    END IF;
    IF EXISTS (SELECT 1 FROM landing_packages_content WHERE gold_landing_package_detail_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing package detail with existing landing packages content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_gold_landing_package_detail_id_fkey';
    END IF;
    IF EXISTS (SELECT 1 FROM landing_packages_content WHERE platinum_landing_package_detail_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing package detail with existing landing packages content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_platinum_landing_package_detail_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_delete_landing_section_header_if_landing_affiliates_con(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_section_header_if_landing_affiliates_con() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_affiliates_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing affiliates content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_affiliates_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_delete_landing_section_header_if_landing_faq_content_ex(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_section_header_if_landing_faq_content_ex() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_faq_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing faq content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_faq_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_delete_landing_section_header_if_landing_features_conte(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_section_header_if_landing_features_conte() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_features_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing features content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_features_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_delete_landing_section_header_if_landing_moments_conten(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_section_header_if_landing_moments_conten() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_moments_content WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing moments content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_moments_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_delete_landing_section_header_if_landing_package_detail(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_section_header_if_landing_package_detail() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_package_details WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_details_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_delete_landing_section_header_if_landing_single_package(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_delete_landing_section_header_if_landing_single_package() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM landing_single_package WHERE landing_section_header_id = OLD.id) THEN
        RAISE EXCEPTION 'Cannot delete landing section header with existing landing single package content'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_content_landing_section_header_id_fkey';
    END IF;
    RETURN OLD;
END;
$$;


--
-- Name: prevent_insert_addon_if_category_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_addon_if_category_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM addon_categories WHERE id = NEW.category_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert addon with soft deleted category'
            USING ERRCODE = '23503', CONSTRAINT = 'addons_category_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_airline_if_logo_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_airline_if_logo_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.logo_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.logo_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert airline with soft deleted logo'
                USING ERRCODE = '23503', CONSTRAINT = 'airlines_logo_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_flight_if_airline_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_flight_if_airline_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM airlines WHERE id = NEW.airline_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight with soft deleted airline'
            USING ERRCODE = '23503', CONSTRAINT = 'flight_routes_flight_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_flight_if_arrival_airport_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_flight_if_arrival_airport_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM airports WHERE id = NEW.arrival_airport_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight with soft deleted arrival airport'
            USING ERRCODE = '23503', CONSTRAINT = 'flights_arrival_airport_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_flight_if_departure_airport_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_flight_if_departure_airport_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM airports WHERE id = NEW.departure_airport_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight with soft deleted departure airport'
            USING ERRCODE = '23503', CONSTRAINT = 'flights_departure_airport_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_flight_route_if_flight_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_flight_route_if_flight_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM flights WHERE id = NEW.flight_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert flight route with soft deleted flight'
            USING ERRCODE = '23503', CONSTRAINT = 'flight_routes_flight_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_flight_route_if_next_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_flight_route_if_next_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM flight_routes WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert flight route with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'flights_return_flight_route_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_guide_if_avatar_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_guide_if_avatar_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.avatar_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.avatar_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert guide with soft deleted avatar'
                USING ERRCODE = '23503', CONSTRAINT = 'guides_avatar_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_hotel_image_if_hotel_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_hotel_image_if_hotel_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM hotels WHERE id = NEW.hotel_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert hotel image with soft deleted hotel'
            USING ERRCODE = '23503', CONSTRAINT = 'hotel_images_hotel_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_hotel_image_if_image_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_hotel_image_if_image_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert hotel image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'hotel_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_day_if_next_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_day_if_next_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_days WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary day with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_days_next_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_day_if_widget_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_day_if_widget_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.widget_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widgets WHERE id = NEW.widget_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary day with soft deleted widget'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_days_widget_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_if_day_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_if_day_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM itinerary_days WHERE id = NEW.day_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary with soft deleted day'
            USING ERRCODE = '23503', CONSTRAINT = 'itineraries_day_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_if_next_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_if_next_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itineraries WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'itineraries_next_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_image_if_image_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_image_if_image_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_image_if_package_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_image_if_package_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM itineraries WHERE id = NEW.itinerary_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary image with soft deleted itinerary'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_images_itinerary_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_activity_image_if_image_is_soft(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_activity_image_if_image_is_soft() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget activity image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_activity_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_activity_image_if_itinerary_wid(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_activity_image_if_itinerary_wid() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM itineraries WHERE id = NEW.itinerary_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget activity image with soft deleted itinerary widget activity'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_activity_images_itinerary_widget_activity_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM hotels WHERE id = NEW.hotel_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget hotel with soft deleted hotel'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_hotels_hotel_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_if_activity_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_if_activity_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.activity_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_activities WHERE id = NEW.activity_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted activity'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_activity_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_if_hotel_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_if_hotel_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.hotel_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_hotels WHERE id = NEW.hotel_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted hotel'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_hotel_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_if_information_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_if_information_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.information_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_informations WHERE id = NEW.information_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted information'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_information_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_if_next_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_if_next_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.next_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widgets WHERE id = NEW.next_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted next'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_next_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_if_recommendation_is_soft_delet(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_if_recommendation_is_soft_delet() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.recommendation_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_recommendations WHERE id = NEW.recommendation_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted recommendation'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_recommendation_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_if_transport_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_if_transport_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.transport_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM itinerary_widget_transports WHERE id = NEW.transport_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert itinerary widget with soft deleted transport'
                USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widgets_transport_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_recommendation_image_if_image_i(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_recommendation_image_if_image_i() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget recommendation image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_recommendation_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_itinerary_widget_recommendation_image_if_itinera(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_itinerary_widget_recommendation_image_if_itinera() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM itineraries WHERE id = NEW.itinerary_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert itinerary widget recommendation image with soft deleted itinerary widget recommendation'
            USING ERRCODE = '23503', CONSTRAINT = 'itinerary_widget_recommendation_images_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_affiliates_content_affiliates_if_image_i(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_affiliates_content_affiliates_if_image_i() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.logo_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.logo_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing affiliates content affiliates with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_affiliates_content_affiliates_logo_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_affiliates_content_if_landing_section_he(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_affiliates_content_if_landing_section_he() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing affiliates content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_affiliates_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_faq_content_if_landing_section_header_is(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_faq_content_if_landing_section_header_is() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing faq content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_faq_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_features_content_benefits_if_image_is_so(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_features_content_benefits_if_image_is_so() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.logo_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.logo_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing features content benefits with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_features_content_benefits_logo_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_features_content_if_landing_section_head(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_features_content_if_landing_section_head() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing features content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_features_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_hero_content_if_image_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_hero_content_if_image_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.image_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert landing hero content with soft deleted image'
                USING ERRCODE = '23503', CONSTRAINT = 'landing_hero_content_image_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_moments_content_if_landing_section_heade(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_moments_content_if_landing_section_heade() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing moments content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_moments_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_moments_content_image_if_image_is_soft_d(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_moments_content_image_if_image_is_soft_d() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing moments content image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_moments_content_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_package_detail_if_landing_section_header(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_package_detail_if_landing_section_header() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package detail with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_details_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_package_detail_item_if_landing_package_d(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_package_detail_item_if_landing_package_d() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package detail item with soft deleted landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_detail_items_landing_package_detail_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_package_detail_item_if_landing_package_i(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_package_detail_item_if_landing_package_i() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package detail item with soft deleted landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_detail_items_landing_package_item_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_package_item_if_package_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_package_item_if_package_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM packages WHERE id = NEW.package_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing package item with soft deleted package'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_package_items_package_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_packages_content_if_landing_package_deta(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_packages_content_if_landing_package_deta() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.silver_landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing packages content with soft deleted silver landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_silver_landing_package_detail_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.gold_landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing packages content with soft deleted gold landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_gold_landing_package_detail_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_details WHERE id = NEW.platinum_landing_package_detail_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing packages content with soft deleted platinum landing package detail'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_packages_content_platinum_landing_package_detail_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_single_package_content_if_landing_packag(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_single_package_content_if_landing_packag() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.silver_landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted silver landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_silver_landing_package_item_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.gold_landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted gold landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_gold_landing_package_item_id_fkey';
    END IF;
    IF (SELECT deleted_at FROM landing_package_items WHERE id = NEW.platinum_landing_package_item_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted platinum landing package item'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_platinum_landing_package_item_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_landing_single_package_content_if_landing_sectio(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_landing_single_package_content_if_landing_sectio() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM landing_section_headers WHERE id = NEW.landing_section_header_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert landing single package content with soft deleted landing section header'
            USING ERRCODE = '23503', CONSTRAINT = 'landing_single_package_content_landing_section_header_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_if_thumbnail_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_if_thumbnail_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.thumbnail_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM images WHERE id = NEW.thumbnail_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package with soft deleted thumbnail'
                USING ERRCODE = '23503', CONSTRAINT = 'packages_thumbnail_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_image_if_image_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_image_if_image_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM images WHERE id = NEW.image_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package image with soft deleted image'
            USING ERRCODE = '23503', CONSTRAINT = 'package_images_image_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_image_if_package_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_image_if_package_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM packages WHERE id = NEW.package_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package image with soft deleted package'
            USING ERRCODE = '23503', CONSTRAINT = 'package_images_package_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_guide_if_guide_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_guide_if_guide_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM guides WHERE id = NEW.guide_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session guide with soft deleted guide'
            USING ERRCODE = '23503', CONSTRAINT = 'package_session_guides_guide_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_guide_if_package_session_is_soft(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_guide_if_package_session_is_soft() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM package_sessions WHERE id = NEW.package_session_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session guide with soft deleted package session'
            USING ERRCODE = '23503', CONSTRAINT = 'package_session_guides_package_session_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_if_bus_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_if_bus_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM buses WHERE id = NEW.bus_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session with soft deleted bus'
            USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_bus_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_if_departure_flight_route_is_sof(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_if_departure_flight_route_is_sof() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.departure_flight_route_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM flight_routes WHERE id = NEW.departure_flight_route_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package session with soft deleted departure flight route'
                USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_departure_flight_route_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_if_embarkation_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_if_embarkation_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT deleted_at FROM embarkations WHERE id = NEW.embarkation_id) IS NOT NULL THEN
        RAISE EXCEPTION 'Cannot insert package session with soft deleted embarkation'
            USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_embarkation_id_fkey';
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_if_package_is_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_if_package_is_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.package_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM packages WHERE id = NEW.package_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package session with soft deleted package'
                USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_package_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: prevent_insert_package_session_if_return_flight_route_is_soft_d(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.prevent_insert_package_session_if_return_flight_route_is_soft_d() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.return_flight_route_id IS NOT NULL THEN
        IF (SELECT deleted_at FROM flight_routes WHERE id = NEW.return_flight_route_id) IS NOT NULL THEN
            RAISE EXCEPTION 'Cannot insert package session with soft deleted return flight route'
                USING ERRCODE = '23503', CONSTRAINT = 'package_sessions_return_flight_route_id_fkey';
        END IF;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_activity_id_null_on_activity_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_activity_id_null_on_activity_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET activity_id = NULL WHERE activity_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_airline_logo_id_null_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_airline_logo_id_null_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE airlines SET logo_id = NULL WHERE logo_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_guide_avatar_id_null_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_guide_avatar_id_null_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE guides SET avatar_id = NULL WHERE avatar_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_hotel_id_null_on_hotel_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_hotel_id_null_on_hotel_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET hotel_id = NULL WHERE hotel_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_information_id_null_on_information_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_information_id_null_on_information_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET information_id = NULL WHERE information_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_landing_affiliates_content_affiliates_logo_id_null_on_image(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_landing_affiliates_content_affiliates_logo_id_null_on_image() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_affiliates_content_affiliates SET logo_id = NULL WHERE logo_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_landing_features_content_benefits_logo_id_null_on_image_sof(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_landing_features_content_benefits_logo_id_null_on_image_sof() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_features_content_benefits SET logo_id = NULL WHERE logo_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_landing_hero_content_image_id_null_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_landing_hero_content_image_id_null_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE landing_hero_content SET image_id = NULL WHERE image_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_landing_single_package_content_landing_package_item_id_null(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_landing_single_package_content_landing_package_item_id_null() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF OLD.deleted_at IS NOT NULL THEN
        UPDATE landing_single_package_content SET silver_landing_package_item_id = NULL WHERE silver_landing_package_item_id = OLD.id;
        UPDATE landing_single_package_content SET gold_landing_package_item_id = NULL WHERE gold_landing_package_item_id = OLD.id;
        UPDATE landing_single_package_content SET platinum_landing_package_item_id = NULL WHERE platinum_landing_package_item_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_next_id_null_on_flight_route_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_next_id_null_on_flight_route_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE flight_routes SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_next_id_null_on_itinerary_day_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_next_id_null_on_itinerary_day_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_days SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_next_id_null_on_itinerary_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_next_id_null_on_itinerary_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itineraries SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_next_id_null_on_itinerary_widget_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_next_id_null_on_itinerary_widget_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET next_id = NULL WHERE next_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_package_session_departure_flight_route_id_null_on_flight_ro(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_package_session_departure_flight_route_id_null_on_flight_ro() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET departure_flight_route_id = NULL WHERE departure_flight_route_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_package_session_return_flight_route_id_null_on_flight_route(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_package_session_return_flight_route_id_null_on_flight_route() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE package_sessions SET return_flight_route_id = NULL WHERE return_flight_route_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_package_thumbnail_id_null_on_image_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_package_thumbnail_id_null_on_image_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE packages SET thumbnail_id = NULL WHERE thumbnail_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_recommendation_id_null_on_recommendation_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_recommendation_id_null_on_recommendation_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET recommendation_id = NULL WHERE recommendation_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_transport_id_null_on_transport_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_transport_id_null_on_transport_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_widgets SET transport_id = NULL WHERE transport_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


--
-- Name: set_widget_id_null_on_widget_soft_deleted(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_widget_id_null_on_widget_soft_deleted() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE itinerary_days SET widget_id = NULL WHERE widget_id = OLD.id;
    END IF;
    RETURN NEW;
END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addon_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.addon_categories (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: addon_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.addon_categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.addon_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: addons; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.addons (
    id bigint NOT NULL,
    category_id bigint NOT NULL,
    name character varying(100) NOT NULL,
    price numeric(13,2) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: addons_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.addons ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.addons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: airlines; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.airlines (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    skytrax_type public.skytrax_type NOT NULL,
    skytrax_rating public.rating NOT NULL,
    logo_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: airlines_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.airlines ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.airlines_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: airports; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.airports (
    id bigint NOT NULL,
    city character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    code character(3) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: airports_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.airports ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.airports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: buses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.buses (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    seat integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    class public.bus_class DEFAULT 'Economy'::public.bus_class NOT NULL
);


--
-- Name: buses_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.buses ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.buses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: city_tours; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.city_tours (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    city character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: city_tours_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.city_tours ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.city_tours_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: embarkations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.embarkations (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    slug character varying(105) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: embarkations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.embarkations ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.embarkations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: facilities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.facilities (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    icon character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: facilities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.facilities ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.facilities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: flight_routes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.flight_routes (
    id bigint NOT NULL,
    flight_id bigint NOT NULL,
    next_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: flight_routes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.flight_routes ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.flight_routes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: flights; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.flights (
    id bigint NOT NULL,
    airline_id bigint NOT NULL,
    aircraft character varying(100) NOT NULL,
    baggage numeric(8,2) NOT NULL,
    cabin_baggage numeric(8,2) NOT NULL,
    departure_airport_id bigint NOT NULL,
    departure_terminal character varying(100),
    departure_at time without time zone NOT NULL,
    arrival_airport_id bigint NOT NULL,
    arrival_terminal character varying(100),
    arrival_at time without time zone NOT NULL,
    code character varying(10) NOT NULL,
    seat_layout character varying(10) NOT NULL,
    class public.flight_class NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: flights_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.flights ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.flights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: guides; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.guides (
    id bigint NOT NULL,
    avatar_id bigint,
    name character varying(100) NOT NULL,
    type public.guide_type NOT NULL,
    description character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: guides_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.guides ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.guides_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: hotel_images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hotel_images (
    hotel_id bigint NOT NULL,
    image_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: hotels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hotels (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    rating public.rating NOT NULL,
    map text NOT NULL,
    address character varying(500) NOT NULL,
    distance numeric(6,2) NOT NULL,
    review text NOT NULL,
    description character varying(1000) NOT NULL,
    location character varying(100) NOT NULL,
    slug character varying(105) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    distance_landmark character varying(100) DEFAULT 'Masjidil Haram'::character varying NOT NULL
);


--
-- Name: hotels_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.hotels ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.hotels_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.images (
    id bigint NOT NULL,
    src character varying(100) NOT NULL,
    alt character varying(100) NOT NULL,
    category character varying(100) DEFAULT NULL::character varying,
    title character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.images ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itineraries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itineraries (
    id bigint NOT NULL,
    city character varying(100) NOT NULL,
    day_id bigint NOT NULL,
    next_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itineraries_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itineraries ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itineraries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_days; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_days (
    id bigint NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    widget_id bigint,
    next_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_days_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_days ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_days_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_images (
    itinerary_id bigint NOT NULL,
    image_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_activities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_activities (
    id bigint NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_widget_activities ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_widget_activities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_widget_activity_images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_activity_images (
    itinerary_widget_activity_id bigint NOT NULL,
    image_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_hotels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_hotels (
    id bigint NOT NULL,
    hotel_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_hotels_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_widget_hotels ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_widget_hotels_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_widget_informations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_informations (
    id bigint NOT NULL,
    description character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_informations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_widget_informations ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_widget_informations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_widget_recommendation_images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_recommendation_images (
    itinerary_widget_recommendation_id bigint NOT NULL,
    image_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_recommendations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_recommendations (
    id bigint NOT NULL,
    description character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_recommendations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_widget_recommendations ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_widget_recommendations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_widget_transports; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widget_transports (
    id bigint NOT NULL,
    transportation character varying(100) NOT NULL,
    "from" character varying(100) NOT NULL,
    "to" character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widget_transports_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_widget_transports ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_widget_transports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: itinerary_widgets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.itinerary_widgets (
    id bigint NOT NULL,
    activity_id bigint,
    hotel_id bigint,
    information_id bigint,
    transport_id bigint,
    recommendation_id bigint,
    next_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: itinerary_widgets_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.itinerary_widgets ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.itinerary_widgets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_affiliates_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_affiliates_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    landing_section_header_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_affiliates_content_id_check CHECK ((id = 1))
);


--
-- Name: landing_affiliates_content_affiliates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_affiliates_content_affiliates (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    name character varying(100) NOT NULL,
    logo_id bigint,
    width integer NOT NULL,
    height integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_affiliates_content_affiliates_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_affiliates_content_affiliates ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_affiliates_content_affiliates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_faq_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_faq_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    landing_section_header_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_faq_content_id_check CHECK ((id = 1))
);


--
-- Name: landing_faq_content_faqs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_faq_content_faqs (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    question character varying(100) NOT NULL,
    answer character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_faq_content_faqs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_faq_content_faqs ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_faq_content_faqs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_features_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_features_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    landing_section_header_id bigint NOT NULL,
    footer_title character varying(100) NOT NULL,
    button_about character varying(100) NOT NULL,
    button_package character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_features_content_id_check CHECK ((id = 1))
);


--
-- Name: landing_features_content_benefits; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_features_content_benefits (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    title character varying(100) NOT NULL,
    subtitle character varying(500) NOT NULL,
    logo_id bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_features_content_benefits_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_features_content_benefits ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_features_content_benefits_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_hero_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_hero_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    tags_line character varying(50) NOT NULL,
    button_label character varying(100) NOT NULL,
    button_url text,
    image_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_hero_content_id_check CHECK ((id = 1))
);


--
-- Name: landing_menus; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_menus (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    icon character varying(100) NOT NULL,
    label character varying(100) NOT NULL,
    path character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_menus_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_menus ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_moments_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_moments_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    landing_section_header_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_moments_content_id_check CHECK ((id = 1))
);


--
-- Name: landing_moments_content_images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_moments_content_images (
    is_enabled boolean DEFAULT true NOT NULL,
    image_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_package_detail_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_package_detail_items (
    is_enabled boolean DEFAULT true NOT NULL,
    landing_package_detail_id bigint NOT NULL,
    landing_package_item_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_package_details; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_package_details (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    landing_section_header_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_package_details_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_package_details ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_package_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_package_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_package_items (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    package_id bigint NOT NULL,
    button_label character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_package_items_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_package_items ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_package_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_packages_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_packages_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    silver_landing_package_detail_id bigint NOT NULL,
    gold_landing_package_detail_id bigint NOT NULL,
    platinum_landing_package_detail_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_packages_content_id_check CHECK ((id = 1))
);


--
-- Name: landing_section_headers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_section_headers (
    id bigint NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    title character varying(100) NOT NULL,
    subtitle character varying(100) DEFAULT NULL::character varying,
    tags_line character varying(50) DEFAULT NULL::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: landing_section_headers_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.landing_section_headers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.landing_section_headers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: landing_single_package_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.landing_single_package_content (
    id integer DEFAULT 1 NOT NULL,
    is_enabled boolean DEFAULT true NOT NULL,
    landing_section_header_id bigint NOT NULL,
    silver_landing_package_item_id bigint,
    gold_landing_package_item_id bigint,
    platinum_landing_package_item_id bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT landing_single_package_content_id_check CHECK ((id = 1))
);


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.migrations (
    version character varying(128) NOT NULL
);


--
-- Name: package_images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.package_images (
    package_id bigint NOT NULL,
    image_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: package_session_guides; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.package_session_guides (
    package_session_id bigint NOT NULL,
    guide_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: package_sessions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.package_sessions (
    id bigint NOT NULL,
    package_id bigint NOT NULL,
    embarkation_id bigint NOT NULL,
    departure_date date NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    departure_flight_route_id bigint,
    return_flight_route_id bigint,
    bus_id bigint,
    itinerary_id bigint,
    double_price numeric(10,2) DEFAULT 0 NOT NULL,
    double_final_price numeric(10,2) DEFAULT NULL::numeric,
    triple_price numeric(10,2) DEFAULT 0 NOT NULL,
    triple_final_price numeric(10,2) DEFAULT NULL::numeric,
    quad_price numeric(10,2) DEFAULT 0 NOT NULL,
    quad_final_price numeric(10,2) DEFAULT NULL::numeric,
    infant_price numeric(10,2) DEFAULT NULL::numeric,
    infant_final_price numeric(10,2) DEFAULT NULL::numeric,
    quota integer DEFAULT 0 NOT NULL
);


--
-- Name: package_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.package_sessions ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.package_sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: packages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.packages (
    id bigint NOT NULL,
    thumbnail_id bigint,
    name character varying(100) NOT NULL,
    category public.package_category NOT NULL,
    type public.package_type NOT NULL,
    slug character varying(105) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    fast_train boolean DEFAULT false
);


--
-- Name: packages_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.packages ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.packages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    full_name character varying(100) NOT NULL,
    phone_number character varying(20) NOT NULL,
    email character varying(256) NOT NULL,
    address character varying(500) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: addon_categories addon_categories_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.addon_categories
    ADD CONSTRAINT addon_categories_id_pkey PRIMARY KEY (id);


--
-- Name: addons addons_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.addons
    ADD CONSTRAINT addons_id_pkey PRIMARY KEY (id);


--
-- Name: airlines airlines_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.airlines
    ADD CONSTRAINT airlines_id_pkey PRIMARY KEY (id);


--
-- Name: airports airports_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.airports
    ADD CONSTRAINT airports_id_pkey PRIMARY KEY (id);


--
-- Name: buses buses_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.buses
    ADD CONSTRAINT buses_id_pkey PRIMARY KEY (id);


--
-- Name: city_tours city_tours_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.city_tours
    ADD CONSTRAINT city_tours_id_pkey PRIMARY KEY (id);


--
-- Name: embarkations embarkations_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.embarkations
    ADD CONSTRAINT embarkations_id_pkey PRIMARY KEY (id);


--
-- Name: facilities facilities_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.facilities
    ADD CONSTRAINT facilities_id_pkey PRIMARY KEY (id);


--
-- Name: flight_routes flight_routes_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flight_routes
    ADD CONSTRAINT flight_routes_id_pkey PRIMARY KEY (id);


--
-- Name: flights flights_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_id_pkey PRIMARY KEY (id);


--
-- Name: guides guides_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.guides
    ADD CONSTRAINT guides_id_pkey PRIMARY KEY (id);


--
-- Name: hotel_images hotel_images_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotel_images
    ADD CONSTRAINT hotel_images_id_pkey PRIMARY KEY (hotel_id, image_id);


--
-- Name: hotels hotels_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotels
    ADD CONSTRAINT hotels_id_pkey PRIMARY KEY (id);


--
-- Name: images images_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_id_pkey PRIMARY KEY (id);


--
-- Name: itineraries itineraries_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itineraries
    ADD CONSTRAINT itineraries_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_days itinerary_days_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_days
    ADD CONSTRAINT itinerary_days_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_images itinerary_images_itinerary_id_image_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_images
    ADD CONSTRAINT itinerary_images_itinerary_id_image_id_key PRIMARY KEY (itinerary_id, image_id);


--
-- Name: itinerary_widget_activities itinerary_widget_activities_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_activities
    ADD CONSTRAINT itinerary_widget_activities_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_widget_activity_images itinerary_widget_activity_images_itinerary_widget_activity_id_i; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_activity_images
    ADD CONSTRAINT itinerary_widget_activity_images_itinerary_widget_activity_id_i PRIMARY KEY (itinerary_widget_activity_id, image_id);


--
-- Name: itinerary_widget_hotels itinerary_widget_hotels_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_hotels
    ADD CONSTRAINT itinerary_widget_hotels_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_widget_informations itinerary_widget_informations_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_informations
    ADD CONSTRAINT itinerary_widget_informations_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_widget_recommendation_images itinerary_widget_recommendation_images_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_recommendation_images
    ADD CONSTRAINT itinerary_widget_recommendation_images_pkey PRIMARY KEY (itinerary_widget_recommendation_id, image_id);


--
-- Name: itinerary_widget_recommendations itinerary_widget_recommendations_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_recommendations
    ADD CONSTRAINT itinerary_widget_recommendations_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_widget_transports itinerary_widget_transports_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_transports
    ADD CONSTRAINT itinerary_widget_transports_id_pkey PRIMARY KEY (id);


--
-- Name: itinerary_widgets itinerary_widgets_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_id_pkey PRIMARY KEY (id);


--
-- Name: landing_affiliates_content_affiliates landing_affiliates_content_affiliates_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_affiliates_content_affiliates
    ADD CONSTRAINT landing_affiliates_content_affiliates_id_pkey PRIMARY KEY (id);


--
-- Name: landing_affiliates_content landing_affiliates_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_affiliates_content
    ADD CONSTRAINT landing_affiliates_content_id_pkey PRIMARY KEY (id);


--
-- Name: landing_faq_content_faqs landing_faq_content_faqs_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_faq_content_faqs
    ADD CONSTRAINT landing_faq_content_faqs_id_pkey PRIMARY KEY (id);


--
-- Name: landing_faq_content landing_faq_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_faq_content
    ADD CONSTRAINT landing_faq_content_id_pkey PRIMARY KEY (id);


--
-- Name: landing_features_content_benefits landing_features_content_benefits_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_features_content_benefits
    ADD CONSTRAINT landing_features_content_benefits_id_pkey PRIMARY KEY (id);


--
-- Name: landing_features_content landing_features_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_features_content
    ADD CONSTRAINT landing_features_content_id_pkey PRIMARY KEY (id);


--
-- Name: landing_hero_content landing_hero_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_hero_content
    ADD CONSTRAINT landing_hero_content_id_pkey PRIMARY KEY (id);


--
-- Name: landing_menus landing_menus_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_menus
    ADD CONSTRAINT landing_menus_id_pkey PRIMARY KEY (id);


--
-- Name: landing_moments_content landing_moments_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_moments_content
    ADD CONSTRAINT landing_moments_content_id_pkey PRIMARY KEY (id);


--
-- Name: landing_moments_content_images landing_moments_content_images_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_moments_content_images
    ADD CONSTRAINT landing_moments_content_images_pkey PRIMARY KEY (image_id);


--
-- Name: landing_package_detail_items landing_package_detail_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_detail_items
    ADD CONSTRAINT landing_package_detail_items_pkey PRIMARY KEY (landing_package_detail_id, landing_package_item_id);


--
-- Name: landing_package_details landing_package_details_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_details
    ADD CONSTRAINT landing_package_details_id_pkey PRIMARY KEY (id);


--
-- Name: landing_package_items landing_package_items_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_items
    ADD CONSTRAINT landing_package_items_id_pkey PRIMARY KEY (id);


--
-- Name: landing_packages_content landing_packages_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_packages_content
    ADD CONSTRAINT landing_packages_content_id_pkey PRIMARY KEY (id);


--
-- Name: landing_section_headers landing_section_headers_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_section_headers
    ADD CONSTRAINT landing_section_headers_id_pkey PRIMARY KEY (id);


--
-- Name: landing_single_package_content landing_single_package_content_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_single_package_content
    ADD CONSTRAINT landing_single_package_content_id_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (version);


--
-- Name: package_images package_images_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_images
    ADD CONSTRAINT package_images_id_pkey PRIMARY KEY (package_id, image_id);


--
-- Name: package_session_guides package_session_guides_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_session_guides
    ADD CONSTRAINT package_session_guides_id_pkey PRIMARY KEY (package_session_id, guide_id);


--
-- Name: package_sessions package_sessions_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_id_pkey PRIMARY KEY (id);


--
-- Name: packages packages_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packages
    ADD CONSTRAINT packages_id_pkey PRIMARY KEY (id);


--
-- Name: users users_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_id_pkey PRIMARY KEY (id);


--
-- Name: addon_categories_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX addon_categories_name_unique ON public.addon_categories USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: addons_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX addons_name_unique ON public.addons USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: airlines_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX airlines_name_unique ON public.airlines USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: airports_code_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX airports_code_unique ON public.airports USING btree (upper((code)::text)) WHERE (deleted_at IS NULL);


--
-- Name: airports_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX airports_name_unique ON public.airports USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: buses_name_class_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX buses_name_class_unique ON public.buses USING btree (upper((name)::text), class) WHERE (deleted_at IS NULL);


--
-- Name: city_tours_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX city_tours_name_unique ON public.city_tours USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: embarkations_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX embarkations_name_unique ON public.embarkations USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: embarkations_slug_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX embarkations_slug_unique ON public.embarkations USING btree (upper((slug)::text)) WHERE (deleted_at IS NULL);


--
-- Name: facilities_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX facilities_name_unique ON public.facilities USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: guides_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX guides_name_unique ON public.guides USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: hotels_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX hotels_name_unique ON public.hotels USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: hotels_slug_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX hotels_slug_unique ON public.hotels USING btree (upper((slug)::text)) WHERE (deleted_at IS NULL);


--
-- Name: images_src_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX images_src_unique ON public.images USING btree (src) WHERE (deleted_at IS NULL);


--
-- Name: package_sessions_package_id_embarkation_id_departure_date_uniqu; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX package_sessions_package_id_embarkation_id_departure_date_uniqu ON public.package_sessions USING btree (package_id, embarkation_id, departure_date) WHERE (deleted_at IS NULL);


--
-- Name: packages_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX packages_name_unique ON public.packages USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: packages_slug_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX packages_slug_unique ON public.packages USING btree (upper((slug)::text)) WHERE (deleted_at IS NULL);


--
-- Name: users_email_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_email_unique ON public.users USING btree (upper((email)::text)) WHERE (deleted_at IS NULL);


--
-- Name: users_phone_number_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_phone_number_unique ON public.users USING btree (upper((phone_number)::text)) WHERE (deleted_at IS NULL);


--
-- Name: addon_categories delete_addon_on_category_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_addon_on_category_soft_deleted BEFORE UPDATE ON public.addon_categories FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_addon_on_category_soft_deleted();


--
-- Name: airlines delete_flight_on_airline_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_flight_on_airline_soft_deleted BEFORE UPDATE ON public.airlines FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_flight_on_airline_soft_deleted();


--
-- Name: airports delete_flight_on_arrival_airport_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_flight_on_arrival_airport_soft_deleted BEFORE UPDATE ON public.airports FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_flight_on_arrival_airport_soft_deleted();


--
-- Name: airports delete_flight_on_departure_airport_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_flight_on_departure_airport_soft_deleted BEFORE UPDATE ON public.airports FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_flight_on_departure_airport_soft_deleted();


--
-- Name: flights delete_flight_route_on_flight_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_flight_route_on_flight_soft_deleted BEFORE UPDATE ON public.flights FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_flight_route_on_flight_soft_deleted();


--
-- Name: hotels delete_hotel_image_on_hotel_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_hotel_image_on_hotel_soft_deleted BEFORE UPDATE ON public.hotels FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_hotel_image_on_hotel_soft_deleted();


--
-- Name: images delete_hotel_image_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_hotel_image_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_hotel_image_on_image_soft_deleted();


--
-- Name: images delete_itinerary_image_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_image_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_image_on_image_soft_deleted();


--
-- Name: itineraries delete_itinerary_image_on_itinerary_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_image_on_itinerary_soft_deleted BEFORE UPDATE ON public.itineraries FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_image_on_itinerary_soft_deleted();


--
-- Name: itinerary_days delete_itinerary_on_day_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_on_day_soft_deleted BEFORE UPDATE ON public.itinerary_days FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_on_day_soft_deleted();


--
-- Name: images delete_itinerary_widget_activity_image_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_widget_activity_image_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_widget_activity_image_on_image_soft_deleted();


--
-- Name: itinerary_widget_activities delete_itinerary_widget_activity_image_on_itinerary_widget_acti; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_widget_activity_image_on_itinerary_widget_acti BEFORE UPDATE ON public.itinerary_widget_activities FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_widget_activity_image_on_itinerary_widget_acti();


--
-- Name: hotels delete_itinerary_widget_hotel_on_hotel_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_widget_hotel_on_hotel_soft_deleted BEFORE UPDATE ON public.hotels FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_widget_hotel_on_hotel_soft_deleted();


--
-- Name: images delete_itinerary_widget_recommendation_image_on_image_soft_dele; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_widget_recommendation_image_on_image_soft_dele BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_widget_recommendation_image_on_image_soft_dele();


--
-- Name: itinerary_widget_recommendations delete_itinerary_widget_recommendation_image_on_itinerary_widge; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_itinerary_widget_recommendation_image_on_itinerary_widge BEFORE UPDATE ON public.itinerary_widget_recommendations FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_itinerary_widget_recommendation_image_on_itinerary_widge();


--
-- Name: images delete_landing_moments_content_image_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_landing_moments_content_image_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_landing_moments_content_image_on_image_soft_deleted();


--
-- Name: landing_package_details delete_landing_package_detail_item_on_landing_package_detail_so; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_landing_package_detail_item_on_landing_package_detail_so BEFORE UPDATE ON public.landing_package_details FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_landing_package_detail_item_on_landing_package_detail_so();


--
-- Name: landing_package_items delete_landing_package_detail_item_on_landing_package_item_soft; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_landing_package_detail_item_on_landing_package_item_soft BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_landing_package_detail_item_on_landing_package_item_soft();


--
-- Name: packages delete_landing_package_item_on_package_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_landing_package_item_on_package_soft_deleted BEFORE UPDATE ON public.packages FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_landing_package_item_on_package_soft_deleted();


--
-- Name: images delete_package_image_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_image_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_image_on_image_soft_deleted();


--
-- Name: packages delete_package_image_on_package_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_image_on_package_soft_deleted BEFORE UPDATE ON public.packages FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_image_on_package_soft_deleted();


--
-- Name: guides delete_package_session_guide_on_guide_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_session_guide_on_guide_soft_deleted BEFORE UPDATE ON public.guides FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_session_guide_on_guide_soft_deleted();


--
-- Name: package_sessions delete_package_session_guide_on_package_session_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_session_guide_on_package_session_soft_deleted BEFORE UPDATE ON public.package_sessions FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_session_guide_on_package_session_soft_deleted();


--
-- Name: buses delete_package_session_on_bus_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_session_on_bus_soft_deleted BEFORE UPDATE ON public.buses FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_session_on_bus_soft_deleted();


--
-- Name: embarkations delete_package_session_on_embarkation_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_session_on_embarkation_soft_deleted BEFORE UPDATE ON public.embarkations FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_session_on_embarkation_soft_deleted();


--
-- Name: packages delete_package_session_on_package_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER delete_package_session_on_package_soft_deleted BEFORE UPDATE ON public.packages FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.delete_package_session_on_package_soft_deleted();


--
-- Name: landing_package_items prevent_delete_landing_package_detail_if_landing_packages_conte; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_package_detail_if_landing_packages_conte BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_package_detail_if_landing_packages_conte();


--
-- Name: landing_package_items prevent_delete_landing_section_header_if_landing_affiliates_con; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_affiliates_con BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_section_header_if_landing_affiliates_con();


--
-- Name: landing_package_items prevent_delete_landing_section_header_if_landing_faq_content_ex; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_faq_content_ex BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_section_header_if_landing_faq_content_ex();


--
-- Name: landing_package_items prevent_delete_landing_section_header_if_landing_features_conte; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_features_conte BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_section_header_if_landing_features_conte();


--
-- Name: landing_package_items prevent_delete_landing_section_header_if_landing_moments_conten; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_moments_conten BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_section_header_if_landing_moments_conten();


--
-- Name: landing_package_items prevent_delete_landing_section_header_if_landing_package_detail; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_package_detail BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_section_header_if_landing_package_detail();


--
-- Name: landing_package_items prevent_delete_landing_section_header_if_landing_single_package; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_delete_landing_section_header_if_landing_single_package BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.prevent_delete_landing_section_header_if_landing_single_package();


--
-- Name: addons prevent_insert_addon_if_category_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_addon_if_category_is_soft_deleted BEFORE INSERT OR UPDATE ON public.addons FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_addon_if_category_is_soft_deleted();


--
-- Name: airlines prevent_insert_airline_if_logo_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_airline_if_logo_is_soft_deleted BEFORE INSERT OR UPDATE ON public.airlines FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_airline_if_logo_is_soft_deleted();


--
-- Name: flights prevent_insert_flight_if_airline_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_flight_if_airline_is_soft_deleted BEFORE INSERT OR UPDATE ON public.flights FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_flight_if_airline_is_soft_deleted();


--
-- Name: flights prevent_insert_flight_if_arrival_airport_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_flight_if_arrival_airport_is_soft_deleted BEFORE INSERT OR UPDATE ON public.flights FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_flight_if_arrival_airport_is_soft_deleted();


--
-- Name: flights prevent_insert_flight_if_departure_airport_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_flight_if_departure_airport_is_soft_deleted BEFORE INSERT OR UPDATE ON public.flights FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_flight_if_departure_airport_is_soft_deleted();


--
-- Name: flight_routes prevent_insert_flight_route_if_flight_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_flight_route_if_flight_is_soft_deleted BEFORE INSERT OR UPDATE ON public.flight_routes FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_flight_route_if_flight_is_soft_deleted();


--
-- Name: flight_routes prevent_insert_flight_route_if_next_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_flight_route_if_next_is_soft_deleted BEFORE INSERT OR UPDATE ON public.flight_routes FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_flight_route_if_next_is_soft_deleted();


--
-- Name: guides prevent_insert_guide_if_avatar_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_guide_if_avatar_is_soft_deleted BEFORE INSERT OR UPDATE ON public.guides FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_guide_if_avatar_is_soft_deleted();


--
-- Name: hotel_images prevent_insert_hotel_image_if_hotel_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_hotel_image_if_hotel_is_soft_deleted BEFORE INSERT OR UPDATE ON public.hotel_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_hotel_image_if_hotel_is_soft_deleted();


--
-- Name: hotel_images prevent_insert_hotel_image_if_image_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_hotel_image_if_image_is_soft_deleted BEFORE INSERT OR UPDATE ON public.hotel_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_hotel_image_if_image_is_soft_deleted();


--
-- Name: itinerary_days prevent_insert_itinerary_day_if_next_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_day_if_next_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_days FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_day_if_next_is_soft_deleted();


--
-- Name: itinerary_days prevent_insert_itinerary_day_if_widget_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_day_if_widget_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_days FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_day_if_widget_is_soft_deleted();


--
-- Name: itineraries prevent_insert_itinerary_if_day_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_if_day_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itineraries FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_if_day_is_soft_deleted();


--
-- Name: itineraries prevent_insert_itinerary_if_next_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_if_next_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itineraries FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_if_next_is_soft_deleted();


--
-- Name: itinerary_images prevent_insert_itinerary_image_if_image_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_image_if_image_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_image_if_image_is_soft_deleted();


--
-- Name: itinerary_images prevent_insert_itinerary_image_if_package_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_image_if_package_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_image_if_package_is_soft_deleted();


--
-- Name: itinerary_widget_activity_images prevent_insert_itinerary_widget_activity_image_if_image_is_soft; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_activity_image_if_image_is_soft BEFORE INSERT OR UPDATE ON public.itinerary_widget_activity_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_activity_image_if_image_is_soft();


--
-- Name: itinerary_widget_activity_images prevent_insert_itinerary_widget_activity_image_if_itinerary_wid; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_activity_image_if_itinerary_wid BEFORE INSERT OR UPDATE ON public.itinerary_widget_activity_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_activity_image_if_itinerary_wid();


--
-- Name: itinerary_widget_hotels prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_widget_hotels FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_hotel_if_hotel_is_soft_deleted();


--
-- Name: itinerary_widgets prevent_insert_itinerary_widget_if_activity_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_if_activity_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_widgets FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_if_activity_is_soft_deleted();


--
-- Name: itinerary_widgets prevent_insert_itinerary_widget_if_hotel_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_if_hotel_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_widgets FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_if_hotel_is_soft_deleted();


--
-- Name: itinerary_widgets prevent_insert_itinerary_widget_if_information_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_if_information_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_widgets FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_if_information_is_soft_deleted();


--
-- Name: itinerary_widgets prevent_insert_itinerary_widget_if_next_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_if_next_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_widgets FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_if_next_is_soft_deleted();


--
-- Name: itinerary_widgets prevent_insert_itinerary_widget_if_recommendation_is_soft_delet; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_if_recommendation_is_soft_delet BEFORE INSERT OR UPDATE ON public.itinerary_widgets FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_if_recommendation_is_soft_delet();


--
-- Name: itinerary_widgets prevent_insert_itinerary_widget_if_transport_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_if_transport_is_soft_deleted BEFORE INSERT OR UPDATE ON public.itinerary_widgets FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_if_transport_is_soft_deleted();


--
-- Name: itinerary_widget_recommendation_images prevent_insert_itinerary_widget_recommendation_image_if_image_i; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_recommendation_image_if_image_i BEFORE INSERT OR UPDATE ON public.itinerary_widget_recommendation_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_recommendation_image_if_image_i();


--
-- Name: itinerary_widget_recommendation_images prevent_insert_itinerary_widget_recommendation_image_if_itinera; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_itinerary_widget_recommendation_image_if_itinera BEFORE INSERT OR UPDATE ON public.itinerary_widget_recommendation_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_itinerary_widget_recommendation_image_if_itinera();


--
-- Name: landing_affiliates_content_affiliates prevent_insert_landing_affiliates_content_affiliates_if_image_i; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_affiliates_content_affiliates_if_image_i BEFORE INSERT OR UPDATE ON public.landing_affiliates_content_affiliates FOR EACH ROW WHEN ((new.logo_id IS NOT NULL)) EXECUTE FUNCTION public.prevent_insert_landing_affiliates_content_affiliates_if_image_i();


--
-- Name: landing_affiliates_content prevent_insert_landing_affiliates_content_if_landing_section_he; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_affiliates_content_if_landing_section_he BEFORE INSERT OR UPDATE ON public.landing_affiliates_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_affiliates_content_if_landing_section_he();


--
-- Name: landing_faq_content prevent_insert_landing_faq_content_if_landing_section_header_is; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_faq_content_if_landing_section_header_is BEFORE INSERT OR UPDATE ON public.landing_faq_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_faq_content_if_landing_section_header_is();


--
-- Name: landing_features_content_benefits prevent_insert_landing_features_content_benefits_if_image_is_so; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_features_content_benefits_if_image_is_so BEFORE INSERT OR UPDATE ON public.landing_features_content_benefits FOR EACH ROW WHEN ((new.logo_id IS NOT NULL)) EXECUTE FUNCTION public.prevent_insert_landing_features_content_benefits_if_image_is_so();


--
-- Name: landing_features_content prevent_insert_landing_features_content_if_landing_section_head; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_features_content_if_landing_section_head BEFORE INSERT OR UPDATE ON public.landing_features_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_features_content_if_landing_section_head();


--
-- Name: landing_hero_content prevent_insert_landing_hero_content_if_image_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_hero_content_if_image_is_soft_deleted BEFORE INSERT OR UPDATE ON public.landing_hero_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_hero_content_if_image_is_soft_deleted();


--
-- Name: landing_moments_content prevent_insert_landing_moments_content_if_landing_section_heade; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_moments_content_if_landing_section_heade BEFORE INSERT OR UPDATE ON public.landing_moments_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_moments_content_if_landing_section_heade();


--
-- Name: landing_moments_content_images prevent_insert_landing_moments_content_image_if_image_is_soft_d; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_moments_content_image_if_image_is_soft_d BEFORE INSERT OR UPDATE ON public.landing_moments_content_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_moments_content_image_if_image_is_soft_d();


--
-- Name: landing_package_details prevent_insert_landing_package_detail_if_landing_section_header; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_package_detail_if_landing_section_header BEFORE INSERT OR UPDATE ON public.landing_package_details FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_package_detail_if_landing_section_header();


--
-- Name: landing_package_detail_items prevent_insert_landing_package_detail_item_if_landing_package_d; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_package_detail_item_if_landing_package_d BEFORE INSERT OR UPDATE ON public.landing_package_detail_items FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_package_detail_item_if_landing_package_d();


--
-- Name: landing_package_detail_items prevent_insert_landing_package_detail_item_if_landing_package_i; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_package_detail_item_if_landing_package_i BEFORE INSERT OR UPDATE ON public.landing_package_detail_items FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_package_detail_item_if_landing_package_i();


--
-- Name: landing_package_items prevent_insert_landing_package_item_if_package_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_package_item_if_package_is_soft_deleted BEFORE INSERT OR UPDATE ON public.landing_package_items FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_package_item_if_package_is_soft_deleted();


--
-- Name: landing_packages_content prevent_insert_landing_packages_content_if_landing_package_deta; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_packages_content_if_landing_package_deta BEFORE INSERT OR UPDATE ON public.landing_packages_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_packages_content_if_landing_package_deta();


--
-- Name: landing_single_package_content prevent_insert_landing_single_package_content_if_landing_packag; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_single_package_content_if_landing_packag BEFORE INSERT OR UPDATE ON public.landing_single_package_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_single_package_content_if_landing_packag();


--
-- Name: landing_single_package_content prevent_insert_landing_single_package_content_if_landing_sectio; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_landing_single_package_content_if_landing_sectio BEFORE INSERT OR UPDATE ON public.landing_single_package_content FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_landing_single_package_content_if_landing_sectio();


--
-- Name: packages prevent_insert_package_if_thumbnail_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_if_thumbnail_is_soft_deleted BEFORE INSERT OR UPDATE ON public.packages FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_if_thumbnail_is_soft_deleted();


--
-- Name: package_images prevent_insert_package_image_if_image_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_image_if_image_is_soft_deleted BEFORE INSERT OR UPDATE ON public.package_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_image_if_image_is_soft_deleted();


--
-- Name: package_images prevent_insert_package_image_if_package_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_image_if_package_is_soft_deleted BEFORE INSERT OR UPDATE ON public.package_images FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_image_if_package_is_soft_deleted();


--
-- Name: package_session_guides prevent_insert_package_session_guide_if_guide_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_guide_if_guide_is_soft_deleted BEFORE INSERT OR UPDATE ON public.package_session_guides FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_guide_if_guide_is_soft_deleted();


--
-- Name: package_session_guides prevent_insert_package_session_guide_if_package_session_is_soft; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_guide_if_package_session_is_soft BEFORE INSERT OR UPDATE ON public.package_session_guides FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_guide_if_package_session_is_soft();


--
-- Name: package_sessions prevent_insert_package_session_if_bus_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_if_bus_is_soft_deleted BEFORE INSERT OR UPDATE ON public.package_sessions FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_if_bus_is_soft_deleted();


--
-- Name: package_sessions prevent_insert_package_session_if_departure_flight_route_is_sof; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_if_departure_flight_route_is_sof BEFORE INSERT OR UPDATE ON public.package_sessions FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_if_departure_flight_route_is_sof();


--
-- Name: package_sessions prevent_insert_package_session_if_embarkation_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_if_embarkation_is_soft_deleted BEFORE INSERT OR UPDATE ON public.package_sessions FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_if_embarkation_is_soft_deleted();


--
-- Name: package_sessions prevent_insert_package_session_if_package_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_if_package_is_soft_deleted BEFORE INSERT OR UPDATE ON public.package_sessions FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_if_package_is_soft_deleted();


--
-- Name: package_sessions prevent_insert_package_session_if_return_flight_route_is_soft_d; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_session_if_return_flight_route_is_soft_d BEFORE INSERT OR UPDATE ON public.package_sessions FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_session_if_return_flight_route_is_soft_d();


--
-- Name: itinerary_widget_activities set_activity_id_null_on_activity_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_activity_id_null_on_activity_soft_deleted BEFORE UPDATE ON public.itinerary_widget_activities FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_activity_id_null_on_activity_soft_deleted();


--
-- Name: images set_airline_logo_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_airline_logo_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_airline_logo_id_null_on_image_soft_deleted();


--
-- Name: images set_guide_avatar_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_guide_avatar_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_guide_avatar_id_null_on_image_soft_deleted();


--
-- Name: itinerary_widget_hotels set_hotel_id_null_on_hotel_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_hotel_id_null_on_hotel_soft_deleted BEFORE UPDATE ON public.itinerary_widget_hotels FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_hotel_id_null_on_hotel_soft_deleted();


--
-- Name: itinerary_widget_informations set_information_id_null_on_information_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_information_id_null_on_information_soft_deleted BEFORE UPDATE ON public.itinerary_widget_informations FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_information_id_null_on_information_soft_deleted();


--
-- Name: images set_landing_affiliates_content_affiliates_logo_id_null_on_image; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_landing_affiliates_content_affiliates_logo_id_null_on_image BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_landing_affiliates_content_affiliates_logo_id_null_on_image();


--
-- Name: images set_landing_features_content_benefits_logo_id_null_on_image_sof; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_landing_features_content_benefits_logo_id_null_on_image_sof BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_landing_features_content_benefits_logo_id_null_on_image_sof();


--
-- Name: images set_landing_hero_content_image_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_landing_hero_content_image_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_landing_hero_content_image_id_null_on_image_soft_deleted();


--
-- Name: landing_package_items set_landing_single_package_content_landing_package_item_id_null; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_landing_single_package_content_landing_package_item_id_null BEFORE UPDATE ON public.landing_package_items FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_landing_single_package_content_landing_package_item_id_null();


--
-- Name: flight_routes set_next_id_null_on_flight_route_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_next_id_null_on_flight_route_soft_deleted BEFORE UPDATE ON public.flight_routes FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_next_id_null_on_flight_route_soft_deleted();


--
-- Name: itinerary_days set_next_id_null_on_itinerary_day_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_next_id_null_on_itinerary_day_soft_deleted BEFORE UPDATE ON public.itinerary_days FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_next_id_null_on_itinerary_day_soft_deleted();


--
-- Name: itineraries set_next_id_null_on_itinerary_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_next_id_null_on_itinerary_soft_deleted BEFORE UPDATE ON public.itineraries FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_next_id_null_on_itinerary_soft_deleted();


--
-- Name: itinerary_widgets set_next_id_null_on_itinerary_widget_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_next_id_null_on_itinerary_widget_soft_deleted BEFORE UPDATE ON public.itinerary_widgets FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_next_id_null_on_itinerary_widget_soft_deleted();


--
-- Name: flight_routes set_package_session_departure_flight_route_id_null_on_flight_ro; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_package_session_departure_flight_route_id_null_on_flight_ro BEFORE UPDATE ON public.flight_routes FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_package_session_departure_flight_route_id_null_on_flight_ro();


--
-- Name: flight_routes set_package_session_return_flight_route_id_null_on_flight_route; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_package_session_return_flight_route_id_null_on_flight_route BEFORE UPDATE ON public.flight_routes FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_package_session_return_flight_route_id_null_on_flight_route();


--
-- Name: images set_package_thumbnail_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_package_thumbnail_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_package_thumbnail_id_null_on_image_soft_deleted();


--
-- Name: itinerary_widget_recommendations set_recommendation_id_null_on_recommendation_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_recommendation_id_null_on_recommendation_soft_deleted BEFORE UPDATE ON public.itinerary_widget_recommendations FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_recommendation_id_null_on_recommendation_soft_deleted();


--
-- Name: itinerary_widget_transports set_transport_id_null_on_transport_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_transport_id_null_on_transport_soft_deleted BEFORE UPDATE ON public.itinerary_widget_transports FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_transport_id_null_on_transport_soft_deleted();


--
-- Name: itinerary_widgets set_widget_id_null_on_widget_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_widget_id_null_on_widget_soft_deleted BEFORE UPDATE ON public.itinerary_widgets FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_widget_id_null_on_widget_soft_deleted();


--
-- Name: addons addons_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.addons
    ADD CONSTRAINT addons_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.addon_categories(id);


--
-- Name: airlines airlines_logo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.airlines
    ADD CONSTRAINT airlines_logo_id_fkey FOREIGN KEY (logo_id) REFERENCES public.images(id);


--
-- Name: flight_routes flight_routes_flight_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flight_routes
    ADD CONSTRAINT flight_routes_flight_id_fkey FOREIGN KEY (flight_id) REFERENCES public.flights(id);


--
-- Name: flight_routes flight_routes_next_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flight_routes
    ADD CONSTRAINT flight_routes_next_id_fkey FOREIGN KEY (next_id) REFERENCES public.flight_routes(id);


--
-- Name: flights flights_airline_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_airline_id_fkey FOREIGN KEY (airline_id) REFERENCES public.airlines(id);


--
-- Name: flights flights_arrival_airport_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_arrival_airport_id_fkey FOREIGN KEY (arrival_airport_id) REFERENCES public.airports(id);


--
-- Name: flights flights_departure_airport_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_departure_airport_id_fkey FOREIGN KEY (departure_airport_id) REFERENCES public.airports(id);


--
-- Name: guides guides_avatar_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.guides
    ADD CONSTRAINT guides_avatar_id_fkey FOREIGN KEY (avatar_id) REFERENCES public.images(id);


--
-- Name: hotel_images hotel_images_hotel_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotel_images
    ADD CONSTRAINT hotel_images_hotel_id_fkey FOREIGN KEY (hotel_id) REFERENCES public.hotels(id);


--
-- Name: hotel_images hotel_images_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hotel_images
    ADD CONSTRAINT hotel_images_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: itineraries itineraries_day_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itineraries
    ADD CONSTRAINT itineraries_day_id_fkey FOREIGN KEY (day_id) REFERENCES public.itinerary_days(id);


--
-- Name: itineraries itineraries_next_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itineraries
    ADD CONSTRAINT itineraries_next_id_fkey FOREIGN KEY (next_id) REFERENCES public.itineraries(id);


--
-- Name: itinerary_days itinerary_days_next_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_days
    ADD CONSTRAINT itinerary_days_next_id_fkey FOREIGN KEY (next_id) REFERENCES public.itinerary_days(id);


--
-- Name: itinerary_days itinerary_days_widget_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_days
    ADD CONSTRAINT itinerary_days_widget_id_fkey FOREIGN KEY (widget_id) REFERENCES public.itinerary_widgets(id);


--
-- Name: itinerary_images itinerary_images_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_images
    ADD CONSTRAINT itinerary_images_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: itinerary_images itinerary_images_itinerary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_images
    ADD CONSTRAINT itinerary_images_itinerary_id_fkey FOREIGN KEY (itinerary_id) REFERENCES public.itineraries(id);


--
-- Name: itinerary_widget_activity_images itinerary_widget_activity_images_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_activity_images
    ADD CONSTRAINT itinerary_widget_activity_images_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: itinerary_widget_activity_images itinerary_widget_activity_images_itinerary_widget_activity_id_f; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_activity_images
    ADD CONSTRAINT itinerary_widget_activity_images_itinerary_widget_activity_id_f FOREIGN KEY (itinerary_widget_activity_id) REFERENCES public.itinerary_widget_activities(id);


--
-- Name: itinerary_widget_hotels itinerary_widget_hotels_hotel_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_hotels
    ADD CONSTRAINT itinerary_widget_hotels_hotel_id_fkey FOREIGN KEY (hotel_id) REFERENCES public.hotels(id);


--
-- Name: itinerary_widget_recommendation_images itinerary_widget_recommendation_images_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_recommendation_images
    ADD CONSTRAINT itinerary_widget_recommendation_images_fkey FOREIGN KEY (itinerary_widget_recommendation_id) REFERENCES public.itinerary_widget_recommendations(id);


--
-- Name: itinerary_widget_recommendation_images itinerary_widget_recommendation_images_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widget_recommendation_images
    ADD CONSTRAINT itinerary_widget_recommendation_images_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: itinerary_widgets itinerary_widgets_activity_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_activity_id_fkey FOREIGN KEY (activity_id) REFERENCES public.itinerary_widget_activities(id);


--
-- Name: itinerary_widgets itinerary_widgets_hotel_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_hotel_id_fkey FOREIGN KEY (hotel_id) REFERENCES public.itinerary_widget_hotels(id);


--
-- Name: itinerary_widgets itinerary_widgets_information_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_information_id_fkey FOREIGN KEY (information_id) REFERENCES public.itinerary_widget_informations(id);


--
-- Name: itinerary_widgets itinerary_widgets_next_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_next_id_fkey FOREIGN KEY (next_id) REFERENCES public.itinerary_widgets(id);


--
-- Name: itinerary_widgets itinerary_widgets_recommendation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_recommendation_id_fkey FOREIGN KEY (recommendation_id) REFERENCES public.itinerary_widget_recommendations(id);


--
-- Name: itinerary_widgets itinerary_widgets_transport_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.itinerary_widgets
    ADD CONSTRAINT itinerary_widgets_transport_id_fkey FOREIGN KEY (transport_id) REFERENCES public.itinerary_widget_transports(id);


--
-- Name: landing_affiliates_content_affiliates landing_affiliates_content_affiliates_logo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_affiliates_content_affiliates
    ADD CONSTRAINT landing_affiliates_content_affiliates_logo_id_fkey FOREIGN KEY (logo_id) REFERENCES public.images(id);


--
-- Name: landing_affiliates_content landing_affiliates_content_landing_section_header_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_affiliates_content
    ADD CONSTRAINT landing_affiliates_content_landing_section_header_id_fkey FOREIGN KEY (landing_section_header_id) REFERENCES public.landing_section_headers(id);


--
-- Name: landing_faq_content landing_faq_content_landing_section_header_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_faq_content
    ADD CONSTRAINT landing_faq_content_landing_section_header_id_fkey FOREIGN KEY (landing_section_header_id) REFERENCES public.landing_section_headers(id);


--
-- Name: landing_features_content_benefits landing_features_content_benefits_logo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_features_content_benefits
    ADD CONSTRAINT landing_features_content_benefits_logo_id_fkey FOREIGN KEY (logo_id) REFERENCES public.images(id);


--
-- Name: landing_features_content landing_features_content_landing_section_header_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_features_content
    ADD CONSTRAINT landing_features_content_landing_section_header_id_fkey FOREIGN KEY (landing_section_header_id) REFERENCES public.landing_section_headers(id);


--
-- Name: landing_hero_content landing_hero_content_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_hero_content
    ADD CONSTRAINT landing_hero_content_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: landing_moments_content_images landing_moments_content_images_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_moments_content_images
    ADD CONSTRAINT landing_moments_content_images_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: landing_moments_content landing_moments_content_landing_section_header_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_moments_content
    ADD CONSTRAINT landing_moments_content_landing_section_header_id_fkey FOREIGN KEY (landing_section_header_id) REFERENCES public.landing_section_headers(id);


--
-- Name: landing_package_detail_items landing_package_detail_items_landing_package_detail_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_detail_items
    ADD CONSTRAINT landing_package_detail_items_landing_package_detail_id_fkey FOREIGN KEY (landing_package_detail_id) REFERENCES public.landing_package_details(id);


--
-- Name: landing_package_detail_items landing_package_detail_items_landing_package_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_detail_items
    ADD CONSTRAINT landing_package_detail_items_landing_package_item_id_fkey FOREIGN KEY (landing_package_item_id) REFERENCES public.landing_package_items(id);


--
-- Name: landing_package_details landing_package_details_landing_section_header_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_details
    ADD CONSTRAINT landing_package_details_landing_section_header_id_fkey FOREIGN KEY (landing_section_header_id) REFERENCES public.landing_section_headers(id);


--
-- Name: landing_package_items landing_package_items_package_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_package_items
    ADD CONSTRAINT landing_package_items_package_id_fkey FOREIGN KEY (package_id) REFERENCES public.packages(id);


--
-- Name: landing_packages_content landing_packages_content_gold_landing_package_detail_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_packages_content
    ADD CONSTRAINT landing_packages_content_gold_landing_package_detail_id_fkey FOREIGN KEY (gold_landing_package_detail_id) REFERENCES public.landing_package_details(id);


--
-- Name: landing_packages_content landing_packages_content_platinum_landing_package_detail_id_fke; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_packages_content
    ADD CONSTRAINT landing_packages_content_platinum_landing_package_detail_id_fke FOREIGN KEY (platinum_landing_package_detail_id) REFERENCES public.landing_package_details(id);


--
-- Name: landing_packages_content landing_packages_content_silver_landing_package_detail_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_packages_content
    ADD CONSTRAINT landing_packages_content_silver_landing_package_detail_id_fkey FOREIGN KEY (silver_landing_package_detail_id) REFERENCES public.landing_package_details(id);


--
-- Name: landing_single_package_content landing_single_package_content_gold_landing_package_item_id_fke; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_single_package_content
    ADD CONSTRAINT landing_single_package_content_gold_landing_package_item_id_fke FOREIGN KEY (gold_landing_package_item_id) REFERENCES public.landing_package_items(id);


--
-- Name: landing_single_package_content landing_single_package_content_landing_section_header_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_single_package_content
    ADD CONSTRAINT landing_single_package_content_landing_section_header_id_fkey FOREIGN KEY (landing_section_header_id) REFERENCES public.landing_section_headers(id);


--
-- Name: landing_single_package_content landing_single_package_content_platinum_landing_package_item_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_single_package_content
    ADD CONSTRAINT landing_single_package_content_platinum_landing_package_item_id FOREIGN KEY (platinum_landing_package_item_id) REFERENCES public.landing_package_items(id);


--
-- Name: landing_single_package_content landing_single_package_content_silver_landing_package_item_id_f; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.landing_single_package_content
    ADD CONSTRAINT landing_single_package_content_silver_landing_package_item_id_f FOREIGN KEY (silver_landing_package_item_id) REFERENCES public.landing_package_items(id);


--
-- Name: package_images package_images_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_images
    ADD CONSTRAINT package_images_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id);


--
-- Name: package_images package_images_package_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_images
    ADD CONSTRAINT package_images_package_id_fkey FOREIGN KEY (package_id) REFERENCES public.packages(id);


--
-- Name: package_session_guides package_session_guides_guide_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_session_guides
    ADD CONSTRAINT package_session_guides_guide_id_fkey FOREIGN KEY (guide_id) REFERENCES public.guides(id);


--
-- Name: package_session_guides package_session_guides_package_session_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_session_guides
    ADD CONSTRAINT package_session_guides_package_session_id_fkey FOREIGN KEY (package_session_id) REFERENCES public.package_sessions(id);


--
-- Name: package_sessions package_sessions_bus_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_bus_id_fkey FOREIGN KEY (bus_id) REFERENCES public.buses(id);


--
-- Name: package_sessions package_sessions_departure_flight_route_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_departure_flight_route_id_fkey FOREIGN KEY (departure_flight_route_id) REFERENCES public.flight_routes(id);


--
-- Name: package_sessions package_sessions_embarkation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_embarkation_id_fkey FOREIGN KEY (embarkation_id) REFERENCES public.embarkations(id);


--
-- Name: package_sessions package_sessions_itinerary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_itinerary_id_fkey FOREIGN KEY (itinerary_id) REFERENCES public.itineraries(id);


--
-- Name: package_sessions package_sessions_package_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_package_id_fkey FOREIGN KEY (package_id) REFERENCES public.packages(id);


--
-- Name: package_sessions package_sessions_return_flight_route_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.package_sessions
    ADD CONSTRAINT package_sessions_return_flight_route_id_fkey FOREIGN KEY (return_flight_route_id) REFERENCES public.flight_routes(id);


--
-- Name: packages packages_thumbnail_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packages
    ADD CONSTRAINT packages_thumbnail_id_fkey FOREIGN KEY (thumbnail_id) REFERENCES public.images(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.migrations (version) VALUES
    ('20250221112243'),
    ('20250221122740'),
    ('20250221131220'),
    ('20250221141043'),
    ('20250221142328'),
    ('20250224082808'),
    ('20250224085253'),
    ('20250224092730'),
    ('20250224114328'),
    ('20250224124038'),
    ('20250226143224'),
    ('20250227124450'),
    ('20250227133727'),
    ('20250227135554'),
    ('20250227142557'),
    ('20250303092458'),
    ('20250304061836'),
    ('20250305052727'),
    ('20250305060747'),
    ('20250313123300'),
    ('20250314090911'),
    ('20250319163517'),
    ('20250320052404'),
    ('20250325063543'),
    ('20250325065013'),
    ('20250325070426'),
    ('20250325070655'),
    ('20250325071654'),
    ('20250325074422');
