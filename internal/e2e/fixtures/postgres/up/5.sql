--
-- PostgreSQL database dump
--

-- Dumped from database version 10.19 (Debian 10.19-1.pgdg90+1)
-- Dumped by pg_dump version 14.1

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

SET default_tablespace = '';

--
-- Name: e2e_user_notes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_user_notes (
    id uuid NOT NULL,
    notes character varying(255),
    user_id uuid NOT NULL,
    slug character varying(64) NOT NULL
);


ALTER TABLE public.e2e_user_notes OWNER TO postgres;

--
-- Name: e2e_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_users (
    id uuid NOT NULL,
    username character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.e2e_users OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: e2e_user_notes e2e_user_notes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_user_notes
    ADD CONSTRAINT e2e_user_notes_pkey PRIMARY KEY (id);


--
-- Name: e2e_users e2e_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_users
    ADD CONSTRAINT e2e_users_pkey PRIMARY KEY (id);


--
-- Name: e2e_user_notes_slug_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX e2e_user_notes_slug_idx ON public.e2e_user_notes USING btree (slug);


--
-- Name: e2e_user_notes_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX e2e_user_notes_user_id_idx ON public.e2e_user_notes USING btree (user_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: e2e_user_notes e2e_user_notes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_user_notes
    ADD CONSTRAINT e2e_user_notes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.e2e_users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

