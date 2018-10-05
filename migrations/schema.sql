--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.10
-- Dumped by pg_dump version 10.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.addresses (
    street character varying(255) NOT NULL,
    house_number integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.addresses OWNER TO postgres;

--
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    title character varying(255) NOT NULL,
    user_id integer,
    isbn character varying(50) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    description character varying(100) DEFAULT 'test'::character varying NOT NULL
);


ALTER TABLE public.books OWNER TO postgres;

--
-- Name: callbacks_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.callbacks_users (
    before_s character varying(255) NOT NULL,
    before_c character varying(255) NOT NULL,
    before_u character varying(255) NOT NULL,
    before_d character varying(255) NOT NULL,
    after_s character varying(255) NOT NULL,
    after_c character varying(255) NOT NULL,
    after_u character varying(255) NOT NULL,
    after_d character varying(255) NOT NULL,
    after_f character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.callbacks_users OWNER TO postgres;

--
-- Name: composers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.composers (
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.composers OWNER TO postgres;

--
-- Name: course_codes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course_codes (
    id uuid NOT NULL,
    course_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.course_codes OWNER TO postgres;

--
-- Name: courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.courses (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- Name: good_friends; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.good_friends (
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL
);


ALTER TABLE public.good_friends OWNER TO postgres;

--
-- Name: not_validatable_cars; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.not_validatable_cars (
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.not_validatable_cars OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: songs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.songs (
    id uuid NOT NULL,
    u_id integer,
    title character varying(255) NOT NULL,
    composed_by_id integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.songs OWNER TO postgres;

--
-- Name: taxis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.taxis (
    model character varying(255) NOT NULL,
    user_id integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.taxis OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    name character varying(255) NOT NULL,
    alive boolean,
    birth_date timestamp without time zone,
    bio text,
    price numeric DEFAULT 1.00,
    email character varying(50) DEFAULT 'foo@example.com'::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_addresses (
    user_id integer NOT NULL,
    address_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users_addresses OWNER TO postgres;

--
-- Name: validatable_cars; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.validatable_cars (
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.validatable_cars OWNER TO postgres;

--
-- Name: writers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.writers (
    name character varying(255) NOT NULL,
    book_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.writers OWNER TO postgres;

--
-- Name: course_codes course_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_codes
    ADD CONSTRAINT course_codes_pkey PRIMARY KEY (id);


--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: songs songs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_pkey PRIMARY KEY (id);


--
-- Name: books_description_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX books_description_index ON public.books USING btree (description);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

