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
-- Name: e2e_address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_address (
    id uuid NOT NULL
);


ALTER TABLE public.e2e_address OWNER TO postgres;

--
-- Name: e2e_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_authors (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.e2e_authors OWNER TO postgres;

--
-- Name: e2e_flow; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_flow (
    id uuid NOT NULL
);


ALTER TABLE public.e2e_flow OWNER TO postgres;

--
-- Name: e2e_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_token (
    id uuid NOT NULL,
    token character varying(64) NOT NULL,
    e2e_flow_id uuid NOT NULL,
    e2e_address_id uuid NOT NULL,
    expires_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    issued_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL
);


ALTER TABLE public.e2e_token OWNER TO postgres;

--
-- Name: e2e_user_posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.e2e_user_posts (
    id uuid NOT NULL,
    author_id uuid,
    slug character varying(32) NOT NULL,
    content character varying(255) DEFAULT ''::character varying NOT NULL,
    published boolean DEFAULT false NOT NULL
);


ALTER TABLE public.e2e_user_posts OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: e2e_address e2e_address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_address
    ADD CONSTRAINT e2e_address_pkey PRIMARY KEY (id);


--
-- Name: e2e_flow e2e_flow_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_flow
    ADD CONSTRAINT e2e_flow_pkey PRIMARY KEY (id);


--
-- Name: e2e_token e2e_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_token
    ADD CONSTRAINT e2e_token_pkey PRIMARY KEY (id);


--
-- Name: e2e_user_posts e2e_user_notes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_user_posts
    ADD CONSTRAINT e2e_user_notes_pkey PRIMARY KEY (id);


--
-- Name: e2e_authors e2e_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_authors
    ADD CONSTRAINT e2e_users_pkey PRIMARY KEY (id);


--
-- Name: e2e_token_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX e2e_token_idx ON public.e2e_token USING btree (token);


--
-- Name: e2e_token_uq_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX e2e_token_uq_idx ON public.e2e_token USING btree (token);


--
-- Name: e2e_user_notes_slug_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX e2e_user_notes_slug_idx ON public.e2e_user_posts USING btree (slug);


--
-- Name: e2e_user_notes_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX e2e_user_notes_user_id_idx ON public.e2e_user_posts USING btree (author_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: e2e_token e2e_token_e2e_address_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_token
    ADD CONSTRAINT e2e_token_e2e_address_id_fkey FOREIGN KEY (e2e_address_id) REFERENCES public.e2e_address(id) ON DELETE CASCADE;


--
-- Name: e2e_token e2e_token_e2e_flow_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_token
    ADD CONSTRAINT e2e_token_e2e_flow_id_fkey FOREIGN KEY (e2e_flow_id) REFERENCES public.e2e_flow(id) ON DELETE CASCADE;


--
-- Name: e2e_user_posts e2e_user_notes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.e2e_user_posts
    ADD CONSTRAINT e2e_user_notes_user_id_fkey FOREIGN KEY (author_id) REFERENCES public.e2e_authors(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

