--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8 (Homebrew)
-- Dumped by pg_dump version 14.8 (Homebrew)

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
-- Name: t_music; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_music (
                                id integer NOT NULL,
                                name character varying,
                                profile_id integer NOT NULL
);


ALTER TABLE public.t_music OWNER TO postgres;

--
-- Name: t_music_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_music_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_music_id_seq OWNER TO postgres;

--
-- Name: t_music_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_music_id_seq OWNED BY public.t_music.id;


--
-- Name: t_music_skill; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_music_skill (
                                      id integer NOT NULL,
                                      music_id integer NOT NULL,
                                      skill_id integer NOT NULL
);


ALTER TABLE public.t_music_skill OWNER TO postgres;

--
-- Name: t_music_skill_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_music_skill_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_music_skill_id_seq OWNER TO postgres;

--
-- Name: t_music_skill_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_music_skill_id_seq OWNED BY public.t_music_skill.id;


--
-- Name: t_profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_profile (
                                  id integer NOT NULL,
                                  name character varying,
                                  birthday date
);


ALTER TABLE public.t_profile OWNER TO postgres;

--
-- Name: t_profile_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_profile_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_profile_id_seq OWNER TO postgres;

--
-- Name: t_profile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_profile_id_seq OWNED BY public.t_profile.id;


--
-- Name: t_project; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_project (
                                  id integer NOT NULL,
                                  name character varying,
                                  type_id integer NOT NULL,
                                  profile_id integer NOT NULL
);


ALTER TABLE public.t_project OWNER TO postgres;

--
-- Name: t_project_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_project_id_seq OWNER TO postgres;

--
-- Name: t_project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_project_id_seq OWNED BY public.t_project.id;


--
-- Name: t_project_skill; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_project_skill (
                                        id integer NOT NULL,
                                        project_id integer NOT NULL,
                                        skill_id integer NOT NULL
);


ALTER TABLE public.t_project_skill OWNER TO postgres;

--
-- Name: t_project_skill_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_project_skill_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_project_skill_id_seq OWNER TO postgres;

--
-- Name: t_project_skill_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_project_skill_id_seq OWNED BY public.t_project_skill.id;


--
-- Name: t_project_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_project_type (
                                       id integer NOT NULL,
                                       name character varying
);


ALTER TABLE public.t_project_type OWNER TO postgres;

--
-- Name: t_project_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_project_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_project_type_id_seq OWNER TO postgres;

--
-- Name: t_project_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_project_type_id_seq OWNED BY public.t_project_type.id;


--
-- Name: t_skill; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_skill (
                                id integer NOT NULL,
                                name character varying,
                                tag_id integer NOT NULL
);


ALTER TABLE public.t_skill OWNER TO postgres;

--
-- Name: t_skill_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_skill_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_skill_id_seq OWNER TO postgres;

--
-- Name: t_skill_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_skill_id_seq OWNED BY public.t_skill.id;


--
-- Name: t_skill_tag; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_skill_tag (
                                    id integer NOT NULL,
                                    name character varying
);


ALTER TABLE public.t_skill_tag OWNER TO postgres;

--
-- Name: t_skill_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_skill_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_skill_tag_id_seq OWNER TO postgres;

--
-- Name: t_skill_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_skill_tag_id_seq OWNED BY public.t_skill_tag.id;


--
-- Name: t_write; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_write (
                                id integer NOT NULL,
                                name character varying,
                                profile_id integer NOT NULL
);


ALTER TABLE public.t_write OWNER TO postgres;

--
-- Name: t_write_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_write_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_write_id_seq OWNER TO postgres;

--
-- Name: t_write_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_write_id_seq OWNED BY public.t_write.id;


--
-- Name: t_music id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music ALTER COLUMN id SET DEFAULT nextval('public.t_music_id_seq'::regclass);


--
-- Name: t_music_skill id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music_skill ALTER COLUMN id SET DEFAULT nextval('public.t_music_skill_id_seq'::regclass);


--
-- Name: t_profile id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_profile ALTER COLUMN id SET DEFAULT nextval('public.t_profile_id_seq'::regclass);


--
-- Name: t_project id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project ALTER COLUMN id SET DEFAULT nextval('public.t_project_id_seq'::regclass);


--
-- Name: t_project_skill id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project_skill ALTER COLUMN id SET DEFAULT nextval('public.t_project_skill_id_seq'::regclass);


--
-- Name: t_project_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project_type ALTER COLUMN id SET DEFAULT nextval('public.t_project_type_id_seq'::regclass);


--
-- Name: t_skill id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill ALTER COLUMN id SET DEFAULT nextval('public.t_skill_id_seq'::regclass);


--
-- Name: t_skill_tag id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill_tag ALTER COLUMN id SET DEFAULT nextval('public.t_skill_tag_id_seq'::regclass);


--
-- Name: t_write id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_write ALTER COLUMN id SET DEFAULT nextval('public.t_write_id_seq'::regclass);


--
-- Data for Name: t_music; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_music (id, name, profile_id) FROM stdin;
\.


--
-- Data for Name: t_music_skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_music_skill (id, music_id, skill_id) FROM stdin;
\.


--
-- Data for Name: t_profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_profile (id, name, birthday) FROM stdin;
1	Z	1999-04-20
\.


--
-- Data for Name: t_project; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_project (id, name, type_id, profile_id) FROM stdin;
3	青蛙王子历险记	1	1
6	动作捕捉程序	2	1
5	幻境赛博射手	1	1
7	戈多赛博射手	1	1
\.


--
-- Data for Name: t_project_skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_project_skill (id, project_id, skill_id) FROM stdin;
2	3	11
4	6	12
5	3	3
6	5	9
7	6	9
3	5	12
8	5	27
\.


--
-- Data for Name: t_project_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_project_type (id, name) FROM stdin;
1	Game
2	Web
3	App
4	Plugin
\.


--
-- Data for Name: t_skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_skill (id, name, tag_id) FROM stdin;
1	Python	1
3	GDScript	1
4	Typescript	1
6	Go	1
8	C++	1
9	C#	1
10	Java	1
11	Godot	2
12	Unity	2
13	Django	3
14	FastAPI	3
15	Gin	3
16	SpringWeb	3
19	Postgres	6
20	Mysql	6
21	Premiere Pro	7
22	Kdenlive	7
23	Photoshop	8
24	Asprite	8
25	Pyxel	8
26	Pixelorama	8
27	Blender	9
28	Material Maker	9
29	FL Studios	10
30	Audacity	10
31	Guitar Pro	10
32	Docker	11
33	Jenkins	11
34	GraphQL	11
35	Figma	12
36	Nginx	13
37	Tomcat	13
38	Selenium	14
39	Requests	15
40	Scrapy	16
41	Urllib	16
42	XPath	16
43	Account Guitar	17
44	Piano	17
45	Elestric Guitar	17
46	Ukulele	17
47	Game	18
48	Music	18
49	Book	18
50	Movie	18
51	Anim	18
52	Run	19
18	Electron	5
17	React Native	4
\.


--
-- Data for Name: t_skill_tag; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_skill_tag (id, name) FROM stdin;
1	Programming Language
2	Backend Framework
3	Web Framework
4	Mobile Framework
5	Desktop Framework
6	Database
7	Video Software
8	2D Art Software
9	3D Art Software
10	Audio Software
11	Dev Tool
12	Design Tool
13	Service Tool
14	Test Framework
16	Crawler Technique
17	Instrument
18	Things I like
19	Sport I like
15	Network Framework
\.


--
-- Data for Name: t_write; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_write (id, name, profile_id) FROM stdin;
\.


--
-- Name: t_music_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_music_id_seq', 1, false);


--
-- Name: t_music_skill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_music_skill_id_seq', 1, false);


--
-- Name: t_profile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_profile_id_seq', 1, false);


--
-- Name: t_project_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_project_id_seq', 7, true);


--
-- Name: t_project_skill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_project_skill_id_seq', 8, true);


--
-- Name: t_project_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_project_type_id_seq', 4, true);


--
-- Name: t_skill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_skill_id_seq', 53, true);


--
-- Name: t_skill_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_skill_tag_id_seq', 20, true);


--
-- Name: t_write_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_write_id_seq', 1, false);


--
-- Name: t_music t_music_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music
    ADD CONSTRAINT t_music_pkey PRIMARY KEY (id);


--
-- Name: t_music_skill t_music_skill_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music_skill
    ADD CONSTRAINT t_music_skill_pkey PRIMARY KEY (id);


--
-- Name: t_profile t_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_profile
    ADD CONSTRAINT t_profile_pkey PRIMARY KEY (id);


--
-- Name: t_project t_project_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project
    ADD CONSTRAINT t_project_pkey PRIMARY KEY (id);


--
-- Name: t_project_skill t_project_skill_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project_skill
    ADD CONSTRAINT t_project_skill_pkey PRIMARY KEY (id);


--
-- Name: t_project_type t_project_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project_type
    ADD CONSTRAINT t_project_type_pkey PRIMARY KEY (id);


--
-- Name: t_skill t_skill_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill
    ADD CONSTRAINT t_skill_pkey PRIMARY KEY (id);


--
-- Name: t_skill_tag t_skill_tag_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill_tag
    ADD CONSTRAINT t_skill_tag_pkey PRIMARY KEY (id);


--
-- Name: t_write t_write_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_write
    ADD CONSTRAINT t_write_pkey PRIMARY KEY (id);


--
-- Name: t_music t_music_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music
    ADD CONSTRAINT t_music_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.t_profile(id);


--
-- Name: t_music_skill t_music_skill_music_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music_skill
    ADD CONSTRAINT t_music_skill_music_id_fkey FOREIGN KEY (music_id) REFERENCES public.t_music(id);


--
-- Name: t_music_skill t_music_skill_skill_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_music_skill
    ADD CONSTRAINT t_music_skill_skill_id_fkey FOREIGN KEY (skill_id) REFERENCES public.t_skill(id);


--
-- Name: t_project t_project_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project
    ADD CONSTRAINT t_project_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.t_profile(id);


--
-- Name: t_project_skill t_project_skill_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project_skill
    ADD CONSTRAINT t_project_skill_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.t_project(id);


--
-- Name: t_project_skill t_project_skill_skill_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project_skill
    ADD CONSTRAINT t_project_skill_skill_id_fkey FOREIGN KEY (skill_id) REFERENCES public.t_skill(id);


--
-- Name: t_project t_project_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_project
    ADD CONSTRAINT t_project_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.t_project_type(id);


--
-- Name: t_skill t_skill_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill
    ADD CONSTRAINT t_skill_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.t_skill_tag(id);


--
-- Name: t_write t_write_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_write
    ADD CONSTRAINT t_write_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.t_profile(id);


--
-- PostgreSQL database dump complete
--

