--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

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

SET default_table_access_method = heap;

--
-- Name: estudiantes; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.estudiantes (
    id integer NOT NULL,
    name character(10),
    age smallint,
    active boolean,
    "createAt" timestamp without time zone DEFAULT now() NOT NULL,
    "updateAt" timestamp without time zone
);


ALTER TABLE public.estudiantes OWNER TO golang;

--
-- Name: estudiantes_id_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.estudiantes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.estudiantes_id_seq OWNER TO golang;

--
-- Name: estudiantes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: golang
--

ALTER SEQUENCE public.estudiantes_id_seq OWNED BY public.estudiantes.id;


--
-- Name: estudiantes id; Type: DEFAULT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.estudiantes ALTER COLUMN id SET DEFAULT nextval('public.estudiantes_id_seq'::regclass);


--
-- Name: estudiantes estudiantes_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.estudiantes
    ADD CONSTRAINT estudiantes_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

