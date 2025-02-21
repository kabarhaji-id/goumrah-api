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


SET default_tablespace = '';

SET default_table_access_method = heap;

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
-- Name: airlines airlines_id_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.airlines
    ADD CONSTRAINT airlines_id_pkey PRIMARY KEY (id);


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
-- Name: airlines_name_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX airlines_name_unique ON public.airlines USING btree (upper((name)::text)) WHERE (deleted_at IS NULL);


--
-- Name: images_src_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX images_src_unique ON public.images USING btree (src) WHERE (deleted_at IS NULL);


--
-- Name: airlines airlines_logo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.airlines
    ADD CONSTRAINT airlines_logo_id_fkey FOREIGN KEY (logo_id) REFERENCES public.images(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.migrations (version) VALUES
    ('20250221112243'),
    ('20250221122740');
