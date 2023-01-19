--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)

-- Started on 2023-01-19 13:05:38 WIB

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
-- TOC entry 5 (class 2615 OID 17840)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3516 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS '';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 219 (class 1259 OID 17866)
-- Name: cities_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities_tab (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.cities_tab OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 17865)
-- Name: cities_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cities_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3518 (class 0 OID 0)
-- Dependencies: 218
-- Name: cities_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_tab_id_seq OWNED BY public.cities_tab.id;


--
-- TOC entry 223 (class 1259 OID 17886)
-- Name: games_chance_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.games_chance_tab (
    id integer NOT NULL,
    user_id integer NOT NULL,
    chance integer,
    count integer,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.games_chance_tab OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 17885)
-- Name: games_chance_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.games_chance_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.games_chance_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3519 (class 0 OID 0)
-- Dependencies: 222
-- Name: games_chance_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.games_chance_tab_id_seq OWNED BY public.games_chance_tab.id;


--
-- TOC entry 225 (class 1259 OID 17895)
-- Name: houses_photos_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.houses_photos_tab (
    id integer NOT NULL,
    house_id integer NOT NULL,
    photo_url character varying,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.houses_photos_tab OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 17894)
-- Name: houses_photos_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.houses_photos_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.houses_photos_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3520 (class 0 OID 0)
-- Dependencies: 224
-- Name: houses_photos_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.houses_photos_tab_id_seq OWNED BY public.houses_photos_tab.id;


--
-- TOC entry 227 (class 1259 OID 17906)
-- Name: houses_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.houses_tab (
    id integer NOT NULL,
    name character varying NOT NULL,
    user_id integer NOT NULL,
    price integer NOT NULL,
    description character varying NOT NULL,
    city_id integer NOT NULL,
    max_guest integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.houses_tab OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 17905)
-- Name: houses_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.houses_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.houses_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3521 (class 0 OID 0)
-- Dependencies: 226
-- Name: houses_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.houses_tab_id_seq OWNED BY public.houses_tab.id;


--
-- TOC entry 233 (class 1259 OID 17935)
-- Name: pickup_status_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickup_status_tab (
    id integer NOT NULL,
    status character varying NOT NULL
);


ALTER TABLE public.pickup_status_tab OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 17934)
-- Name: pickup_status_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pickup_status_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pickup_status_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3522 (class 0 OID 0)
-- Dependencies: 232
-- Name: pickup_status_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pickup_status_tab_id_seq OWNED BY public.pickup_status_tab.id;


--
-- TOC entry 231 (class 1259 OID 17926)
-- Name: pickups_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickups_tab (
    id integer NOT NULL,
    user_id integer NOT NULL,
    reservation_id integer NOT NULL,
    pickup_status_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.pickups_tab OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 17925)
-- Name: pickups_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pickups_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pickups_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3523 (class 0 OID 0)
-- Dependencies: 230
-- Name: pickups_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pickups_tab_id_seq OWNED BY public.pickups_tab.id;


--
-- TOC entry 229 (class 1259 OID 17917)
-- Name: reservations_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reservations_tab (
    id integer NOT NULL,
    house_id integer NOT NULL,
    user_id integer NOT NULL,
    check_in_date timestamp without time zone NOT NULL,
    check_out_date timestamp without time zone NOT NULL,
    total_price integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.reservations_tab OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 17916)
-- Name: reservations_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reservations_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reservations_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3524 (class 0 OID 0)
-- Dependencies: 228
-- Name: reservations_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reservations_tab_id_seq OWNED BY public.reservations_tab.id;


--
-- TOC entry 217 (class 1259 OID 17855)
-- Name: roles_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles_tab (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.roles_tab OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 17854)
-- Name: roles_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3525 (class 0 OID 0)
-- Dependencies: 216
-- Name: roles_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_tab_id_seq OWNED BY public.roles_tab.id;


--
-- TOC entry 237 (class 1259 OID 17953)
-- Name: transaction_type_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_type_tab (
    id integer NOT NULL,
    type character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.transaction_type_tab OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 17952)
-- Name: transaction_type_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transaction_type_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transaction_type_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3526 (class 0 OID 0)
-- Dependencies: 236
-- Name: transaction_type_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transaction_type_tab_id_seq OWNED BY public.transaction_type_tab.id;


--
-- TOC entry 235 (class 1259 OID 17944)
-- Name: transactions_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions_tab (
    id integer NOT NULL,
    wallet_id integer NOT NULL,
    transaction_type_id integer NOT NULL,
    balance integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.transactions_tab OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 17943)
-- Name: transactions_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3527 (class 0 OID 0)
-- Dependencies: 234
-- Name: transactions_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_tab_id_seq OWNED BY public.transactions_tab.id;


--
-- TOC entry 215 (class 1259 OID 17842)
-- Name: users_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_tab (
    id integer NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    full_name character varying,
    address character varying,
    city_id integer NOT NULL,
    role_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.users_tab OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 17841)
-- Name: users_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3528 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_tab_id_seq OWNED BY public.users_tab.id;


--
-- TOC entry 221 (class 1259 OID 17877)
-- Name: wallets_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets_tab (
    id integer NOT NULL,
    user_id integer NOT NULL,
    balance integer,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.wallets_tab OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 17876)
-- Name: wallets_tab_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_tab_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_tab_id_seq OWNER TO postgres;

--
-- TOC entry 3529 (class 0 OID 0)
-- Dependencies: 220
-- Name: wallets_tab_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallets_tab_id_seq OWNED BY public.wallets_tab.id;


--
-- TOC entry 3277 (class 2604 OID 17869)
-- Name: cities_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities_tab ALTER COLUMN id SET DEFAULT nextval('public.cities_tab_id_seq'::regclass);


--
-- TOC entry 3283 (class 2604 OID 17889)
-- Name: games_chance_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.games_chance_tab ALTER COLUMN id SET DEFAULT nextval('public.games_chance_tab_id_seq'::regclass);


--
-- TOC entry 3286 (class 2604 OID 17898)
-- Name: houses_photos_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_photos_tab ALTER COLUMN id SET DEFAULT nextval('public.houses_photos_tab_id_seq'::regclass);


--
-- TOC entry 3289 (class 2604 OID 17909)
-- Name: houses_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_tab ALTER COLUMN id SET DEFAULT nextval('public.houses_tab_id_seq'::regclass);


--
-- TOC entry 3298 (class 2604 OID 17938)
-- Name: pickup_status_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickup_status_tab ALTER COLUMN id SET DEFAULT nextval('public.pickup_status_tab_id_seq'::regclass);


--
-- TOC entry 3295 (class 2604 OID 17929)
-- Name: pickups_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_tab ALTER COLUMN id SET DEFAULT nextval('public.pickups_tab_id_seq'::regclass);


--
-- TOC entry 3292 (class 2604 OID 17920)
-- Name: reservations_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations_tab ALTER COLUMN id SET DEFAULT nextval('public.reservations_tab_id_seq'::regclass);


--
-- TOC entry 3274 (class 2604 OID 17858)
-- Name: roles_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles_tab ALTER COLUMN id SET DEFAULT nextval('public.roles_tab_id_seq'::regclass);


--
-- TOC entry 3302 (class 2604 OID 17956)
-- Name: transaction_type_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_type_tab ALTER COLUMN id SET DEFAULT nextval('public.transaction_type_tab_id_seq'::regclass);


--
-- TOC entry 3299 (class 2604 OID 17947)
-- Name: transactions_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_tab ALTER COLUMN id SET DEFAULT nextval('public.transactions_tab_id_seq'::regclass);


--
-- TOC entry 3271 (class 2604 OID 17845)
-- Name: users_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tab ALTER COLUMN id SET DEFAULT nextval('public.users_tab_id_seq'::regclass);


--
-- TOC entry 3280 (class 2604 OID 17880)
-- Name: wallets_tab id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets_tab ALTER COLUMN id SET DEFAULT nextval('public.wallets_tab_id_seq'::regclass);


--
-- TOC entry 3492 (class 0 OID 17866)
-- Dependencies: 219
-- Data for Name: cities_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3496 (class 0 OID 17886)
-- Dependencies: 223
-- Data for Name: games_chance_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3498 (class 0 OID 17895)
-- Dependencies: 225
-- Data for Name: houses_photos_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3500 (class 0 OID 17906)
-- Dependencies: 227
-- Data for Name: houses_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3506 (class 0 OID 17935)
-- Dependencies: 233
-- Data for Name: pickup_status_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3504 (class 0 OID 17926)
-- Dependencies: 231
-- Data for Name: pickups_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3502 (class 0 OID 17917)
-- Dependencies: 229
-- Data for Name: reservations_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3490 (class 0 OID 17855)
-- Dependencies: 217
-- Data for Name: roles_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3510 (class 0 OID 17953)
-- Dependencies: 237
-- Data for Name: transaction_type_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3508 (class 0 OID 17944)
-- Dependencies: 235
-- Data for Name: transactions_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3488 (class 0 OID 17842)
-- Dependencies: 215
-- Data for Name: users_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3494 (class 0 OID 17877)
-- Dependencies: 221
-- Data for Name: wallets_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3530 (class 0 OID 0)
-- Dependencies: 218
-- Name: cities_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_tab_id_seq', 1, false);


--
-- TOC entry 3531 (class 0 OID 0)
-- Dependencies: 222
-- Name: games_chance_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.games_chance_tab_id_seq', 1, false);


--
-- TOC entry 3532 (class 0 OID 0)
-- Dependencies: 224
-- Name: houses_photos_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.houses_photos_tab_id_seq', 1, false);


--
-- TOC entry 3533 (class 0 OID 0)
-- Dependencies: 226
-- Name: houses_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.houses_tab_id_seq', 1, false);


--
-- TOC entry 3534 (class 0 OID 0)
-- Dependencies: 232
-- Name: pickup_status_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pickup_status_tab_id_seq', 1, false);


--
-- TOC entry 3535 (class 0 OID 0)
-- Dependencies: 230
-- Name: pickups_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pickups_tab_id_seq', 1, false);


--
-- TOC entry 3536 (class 0 OID 0)
-- Dependencies: 228
-- Name: reservations_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reservations_tab_id_seq', 1, false);


--
-- TOC entry 3537 (class 0 OID 0)
-- Dependencies: 216
-- Name: roles_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_tab_id_seq', 1, false);


--
-- TOC entry 3538 (class 0 OID 0)
-- Dependencies: 236
-- Name: transaction_type_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transaction_type_tab_id_seq', 1, false);


--
-- TOC entry 3539 (class 0 OID 0)
-- Dependencies: 234
-- Name: transactions_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_tab_id_seq', 1, false);


--
-- TOC entry 3540 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_tab_id_seq', 1, false);


--
-- TOC entry 3541 (class 0 OID 0)
-- Dependencies: 220
-- Name: wallets_tab_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_tab_id_seq', 1, false);


--
-- TOC entry 3312 (class 2606 OID 17875)
-- Name: cities_tab cities_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities_tab
    ADD CONSTRAINT cities_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3316 (class 2606 OID 17893)
-- Name: games_chance_tab games_chance_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.games_chance_tab
    ADD CONSTRAINT games_chance_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3318 (class 2606 OID 17904)
-- Name: houses_photos_tab houses_photos_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_photos_tab
    ADD CONSTRAINT houses_photos_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3320 (class 2606 OID 17915)
-- Name: houses_tab houses_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_tab
    ADD CONSTRAINT houses_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3326 (class 2606 OID 17942)
-- Name: pickup_status_tab pickup_status_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickup_status_tab
    ADD CONSTRAINT pickup_status_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3324 (class 2606 OID 17933)
-- Name: pickups_tab pickups_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_tab
    ADD CONSTRAINT pickups_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3322 (class 2606 OID 17924)
-- Name: reservations_tab reservations_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations_tab
    ADD CONSTRAINT reservations_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3310 (class 2606 OID 17864)
-- Name: roles_tab roles_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles_tab
    ADD CONSTRAINT roles_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3330 (class 2606 OID 17960)
-- Name: transaction_type_tab transaction_type_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_type_tab
    ADD CONSTRAINT transaction_type_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3328 (class 2606 OID 17951)
-- Name: transactions_tab transactions_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_tab
    ADD CONSTRAINT transactions_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3306 (class 2606 OID 17853)
-- Name: users_tab users_tab_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tab
    ADD CONSTRAINT users_tab_email_key UNIQUE (email);


--
-- TOC entry 3308 (class 2606 OID 17851)
-- Name: users_tab users_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tab
    ADD CONSTRAINT users_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3314 (class 2606 OID 17884)
-- Name: wallets_tab wallets_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets_tab
    ADD CONSTRAINT wallets_tab_pkey PRIMARY KEY (id);


--
-- TOC entry 3334 (class 2606 OID 17976)
-- Name: games_chance_tab games_chance_tab_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.games_chance_tab
    ADD CONSTRAINT games_chance_tab_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_tab(id);


--
-- TOC entry 3335 (class 2606 OID 17981)
-- Name: houses_photos_tab houses_photos_tab_house_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_photos_tab
    ADD CONSTRAINT houses_photos_tab_house_id_fkey FOREIGN KEY (house_id) REFERENCES public.houses_tab(id);


--
-- TOC entry 3336 (class 2606 OID 17991)
-- Name: houses_tab houses_tab_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_tab
    ADD CONSTRAINT houses_tab_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities_tab(id);


--
-- TOC entry 3337 (class 2606 OID 17986)
-- Name: houses_tab houses_tab_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_tab
    ADD CONSTRAINT houses_tab_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_tab(id);


--
-- TOC entry 3340 (class 2606 OID 18016)
-- Name: pickups_tab pickups_tab_pickup_status_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_tab
    ADD CONSTRAINT pickups_tab_pickup_status_id_fkey FOREIGN KEY (pickup_status_id) REFERENCES public.pickup_status_tab(id);


--
-- TOC entry 3341 (class 2606 OID 18011)
-- Name: pickups_tab pickups_tab_reservation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_tab
    ADD CONSTRAINT pickups_tab_reservation_id_fkey FOREIGN KEY (reservation_id) REFERENCES public.reservations_tab(id);


--
-- TOC entry 3342 (class 2606 OID 18006)
-- Name: pickups_tab pickups_tab_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_tab
    ADD CONSTRAINT pickups_tab_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_tab(id);


--
-- TOC entry 3338 (class 2606 OID 17996)
-- Name: reservations_tab reservations_tab_house_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations_tab
    ADD CONSTRAINT reservations_tab_house_id_fkey FOREIGN KEY (house_id) REFERENCES public.houses_tab(id);


--
-- TOC entry 3339 (class 2606 OID 18001)
-- Name: reservations_tab reservations_tab_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations_tab
    ADD CONSTRAINT reservations_tab_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_tab(id);


--
-- TOC entry 3343 (class 2606 OID 18026)
-- Name: transactions_tab transactions_tab_transaction_type_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_tab
    ADD CONSTRAINT transactions_tab_transaction_type_fkey FOREIGN KEY (transaction_type_id) REFERENCES public.transaction_type_tab(id);


--
-- TOC entry 3344 (class 2606 OID 18021)
-- Name: transactions_tab transactions_tab_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_tab
    ADD CONSTRAINT transactions_tab_user_id_fkey FOREIGN KEY (wallet_id) REFERENCES public.wallets_tab(id);


--
-- TOC entry 3331 (class 2606 OID 17961)
-- Name: users_tab users_tab_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tab
    ADD CONSTRAINT users_tab_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities_tab(id);


--
-- TOC entry 3332 (class 2606 OID 17966)
-- Name: users_tab users_tab_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tab
    ADD CONSTRAINT users_tab_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles_tab(id);


--
-- TOC entry 3333 (class 2606 OID 17971)
-- Name: wallets_tab wallets_tab_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets_tab
    ADD CONSTRAINT wallets_tab_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_tab(id);


--
-- TOC entry 3517 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-01-19 13:05:38 WIB

--
-- PostgreSQL database dump complete
--

