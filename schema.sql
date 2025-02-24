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
-- Name: migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.migrations (
    version character varying(128) NOT NULL
);


--
-- Name: packages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.packages (
    id bigint NOT NULL,
    thumbnail_id bigint,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    is_active boolean DEFAULT false NOT NULL,
    category public.package_category NOT NULL,
    type public.package_type NOT NULL,
    slug character varying(105) NOT NULL,
    is_recommended boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
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
-- Name: addon_categories addon_categories_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.addon_categories
    ADD CONSTRAINT addon_categories_id_pkey PRIMARY KEY (id);


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
-- Name: embarkations embarkations_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.embarkations
    ADD CONSTRAINT embarkations_id_pkey PRIMARY KEY (id);


--
-- Name: guides guides_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.guides
    ADD CONSTRAINT guides_id_pkey PRIMARY KEY (id);


--
-- Name: images images_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_id_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (version);


--
-- Name: packages packages_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packages
    ADD CONSTRAINT packages_id_pkey PRIMARY KEY (id);


--
-- Name: addon_categories_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX addon_categories_name_unique ON public.addon_categories USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


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
-- Name: embarkations_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX embarkations_name_unique ON public.embarkations USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: embarkations_slug_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX embarkations_slug_unique ON public.embarkations USING btree (upper((slug)::text)) WHERE (deleted_at IS NULL);


--
-- Name: guides_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX guides_name_unique ON public.guides USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: images_src_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX images_src_unique ON public.images USING btree (src) WHERE (deleted_at IS NULL);


--
-- Name: packages_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX packages_name_unique ON public.packages USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: packages_slug_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX packages_slug_unique ON public.packages USING btree (upper((slug)::text)) WHERE (deleted_at IS NULL);


--
-- Name: airlines prevent_insert_airline_if_logo_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_airline_if_logo_is_soft_deleted BEFORE INSERT OR UPDATE ON public.airlines FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_airline_if_logo_is_soft_deleted();


--
-- Name: guides prevent_insert_guide_if_avatar_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_guide_if_avatar_is_soft_deleted BEFORE INSERT OR UPDATE ON public.guides FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_guide_if_avatar_is_soft_deleted();


--
-- Name: packages prevent_insert_package_if_thumbnail_is_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER prevent_insert_package_if_thumbnail_is_soft_deleted BEFORE INSERT OR UPDATE ON public.packages FOR EACH ROW EXECUTE FUNCTION public.prevent_insert_package_if_thumbnail_is_soft_deleted();


--
-- Name: images set_airline_logo_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_airline_logo_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_airline_logo_id_null_on_image_soft_deleted();


--
-- Name: images set_guide_avatar_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_guide_avatar_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_guide_avatar_id_null_on_image_soft_deleted();


--
-- Name: images set_package_thumbnail_id_null_on_image_soft_deleted; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_package_thumbnail_id_null_on_image_soft_deleted BEFORE UPDATE ON public.images FOR EACH ROW WHEN (((old.deleted_at IS NULL) AND (new.deleted_at IS NOT NULL))) EXECUTE FUNCTION public.set_package_thumbnail_id_null_on_image_soft_deleted();


--
-- Name: airlines airlines_logo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.airlines
    ADD CONSTRAINT airlines_logo_id_fkey FOREIGN KEY (logo_id) REFERENCES public.images(id) ON DELETE SET NULL;


--
-- Name: guides guides_avatar_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.guides
    ADD CONSTRAINT guides_avatar_id_fkey FOREIGN KEY (avatar_id) REFERENCES public.images(id);


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
    ('20250224092730');
