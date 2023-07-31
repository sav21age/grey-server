--
-- PostgreSQL database dump
--

-- Dumped from database version 11.19
-- Dumped by pg_dump version 11.19

-- Started on 2023-07-31 19:02:04

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
-- TOC entry 2928 (class 1262 OID 167331)
-- Name: grey; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE grey WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE grey OWNER TO postgres;

\connect grey

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
-- TOC entry 7 (class 2615 OID 167339)
-- Name: grey; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA grey;


ALTER SCHEMA grey OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 208 (class 1259 OID 167412)
-- Name: cart; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.cart (
    id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE grey.cart OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 167410)
-- Name: cart_id_seq; Type: SEQUENCE; Schema: grey; Owner: postgres
--

CREATE SEQUENCE grey.cart_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE grey.cart_id_seq OWNER TO postgres;

--
-- TOC entry 2929 (class 0 OID 0)
-- Dependencies: 207
-- Name: cart_id_seq; Type: SEQUENCE OWNED BY; Schema: grey; Owner: postgres
--

ALTER SEQUENCE grey.cart_id_seq OWNED BY grey.cart.id;


--
-- TOC entry 209 (class 1259 OID 167425)
-- Name: cart_item; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.cart_item (
    product_id integer NOT NULL,
    cart_id integer NOT NULL,
    price_id integer NOT NULL,
    quantity integer,
    CONSTRAINT cart_item_quantity_check CHECK ((quantity > 0))
);


ALTER TABLE grey.cart_item OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 167448)
-- Name: order; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey."order" (
    id integer NOT NULL,
    user_id integer NOT NULL,
    date timestamp without time zone DEFAULT now()
);


ALTER TABLE grey."order" OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 167446)
-- Name: order_id_seq; Type: SEQUENCE; Schema: grey; Owner: postgres
--

CREATE SEQUENCE grey.order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE grey.order_id_seq OWNER TO postgres;

--
-- TOC entry 2930 (class 0 OID 0)
-- Dependencies: 210
-- Name: order_id_seq; Type: SEQUENCE OWNED BY; Schema: grey; Owner: postgres
--

ALTER SEQUENCE grey.order_id_seq OWNED BY grey."order".id;


--
-- TOC entry 212 (class 1259 OID 167459)
-- Name: order_item; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.order_item (
    product_id integer NOT NULL,
    order_id integer NOT NULL,
    price_id integer NOT NULL,
    quantity integer,
    CONSTRAINT order_item_quantity_check CHECK ((quantity > 0))
);


ALTER TABLE grey.order_item OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 167395)
-- Name: price; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.price (
    id integer NOT NULL,
    product_id integer NOT NULL,
    price numeric,
    date timestamp without time zone NOT NULL,
    CONSTRAINT price_price_check CHECK ((price > (0)::numeric))
);


ALTER TABLE grey.price OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 167393)
-- Name: price_id_seq; Type: SEQUENCE; Schema: grey; Owner: postgres
--

CREATE SEQUENCE grey.price_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE grey.price_id_seq OWNER TO postgres;

--
-- TOC entry 2931 (class 0 OID 0)
-- Dependencies: 205
-- Name: price_id_seq; Type: SEQUENCE OWNED BY; Schema: grey; Owner: postgres
--

ALTER SEQUENCE grey.price_id_seq OWNED BY grey.price.id;


--
-- TOC entry 201 (class 1259 OID 167357)
-- Name: product; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.product (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    description text,
    quantity integer NOT NULL,
    price_date timestamp without time zone NOT NULL
);


ALTER TABLE grey.product OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 167355)
-- Name: product_id_seq; Type: SEQUENCE; Schema: grey; Owner: postgres
--

CREATE SEQUENCE grey.product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE grey.product_id_seq OWNER TO postgres;

--
-- TOC entry 2932 (class 0 OID 0)
-- Dependencies: 200
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: grey; Owner: postgres
--

ALTER SEQUENCE grey.product_id_seq OWNED BY grey.product.id;


--
-- TOC entry 204 (class 1259 OID 167378)
-- Name: product_tag; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.product_tag (
    product_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE grey.product_tag OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 167370)
-- Name: tag; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey.tag (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE grey.tag OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 167368)
-- Name: tag_id_seq; Type: SEQUENCE; Schema: grey; Owner: postgres
--

CREATE SEQUENCE grey.tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE grey.tag_id_seq OWNER TO postgres;

--
-- TOC entry 2933 (class 0 OID 0)
-- Dependencies: 202
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: grey; Owner: postgres
--

ALTER SEQUENCE grey.tag_id_seq OWNED BY grey.tag.id;


--
-- TOC entry 199 (class 1259 OID 167342)
-- Name: user; Type: TABLE; Schema: grey; Owner: postgres
--

CREATE TABLE grey."user" (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    firstname character varying(100) NOT NULL,
    lastname character varying(100) NOT NULL,
    fullname character varying(200) NOT NULL,
    age integer NOT NULL,
    is_married boolean DEFAULT false NOT NULL,
    password character varying(255) NOT NULL,
    CONSTRAINT user_age_check CHECK ((age > 17))
);


ALTER TABLE grey."user" OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 167340)
-- Name: user_id_seq; Type: SEQUENCE; Schema: grey; Owner: postgres
--

CREATE SEQUENCE grey.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE grey.user_id_seq OWNER TO postgres;

--
-- TOC entry 2934 (class 0 OID 0)
-- Dependencies: 198
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: grey; Owner: postgres
--

ALTER SEQUENCE grey.user_id_seq OWNED BY grey."user".id;


--
-- TOC entry 197 (class 1259 OID 167332)
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- TOC entry 2742 (class 2604 OID 167415)
-- Name: cart id; Type: DEFAULT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart ALTER COLUMN id SET DEFAULT nextval('grey.cart_id_seq'::regclass);


--
-- TOC entry 2744 (class 2604 OID 167451)
-- Name: order id; Type: DEFAULT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey."order" ALTER COLUMN id SET DEFAULT nextval('grey.order_id_seq'::regclass);


--
-- TOC entry 2740 (class 2604 OID 167398)
-- Name: price id; Type: DEFAULT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.price ALTER COLUMN id SET DEFAULT nextval('grey.price_id_seq'::regclass);


--
-- TOC entry 2738 (class 2604 OID 167360)
-- Name: product id; Type: DEFAULT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.product ALTER COLUMN id SET DEFAULT nextval('grey.product_id_seq'::regclass);


--
-- TOC entry 2739 (class 2604 OID 167373)
-- Name: tag id; Type: DEFAULT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.tag ALTER COLUMN id SET DEFAULT nextval('grey.tag_id_seq'::regclass);


--
-- TOC entry 2735 (class 2604 OID 167345)
-- Name: user id; Type: DEFAULT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey."user" ALTER COLUMN id SET DEFAULT nextval('grey.user_id_seq'::regclass);


--
-- TOC entry 2918 (class 0 OID 167412)
-- Dependencies: 208
-- Data for Name: cart; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.cart (id, user_id) FROM stdin;
13	1
\.


--
-- TOC entry 2919 (class 0 OID 167425)
-- Dependencies: 209
-- Data for Name: cart_item; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.cart_item (product_id, cart_id, price_id, quantity) FROM stdin;
1	13	7	2
2	13	8	2
3	13	10	2
\.


--
-- TOC entry 2921 (class 0 OID 167448)
-- Dependencies: 211
-- Data for Name: order; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey."order" (id, user_id, date) FROM stdin;
11	1	2023-07-31 08:17:25.520355
12	1	2023-07-31 08:17:25.520355
13	2	2023-07-31 08:17:25.520355
\.


--
-- TOC entry 2922 (class 0 OID 167459)
-- Dependencies: 212
-- Data for Name: order_item; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.order_item (product_id, order_id, price_id, quantity) FROM stdin;
1	11	6	2
2	11	2	1
3	11	3	1
3	12	3	5
3	13	10	2
2	13	8	2
1	13	7	2
\.


--
-- TOC entry 2916 (class 0 OID 167395)
-- Dependencies: 206
-- Data for Name: price; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.price (id, product_id, price, date) FROM stdin;
1	1	50	2023-07-30 18:58:18.248426
2	2	75	2023-07-30 18:58:49.622435
3	3	35	2023-07-30 18:59:15.998929
4	1	56	2023-07-30 19:00:32.396135
5	1	61	2023-07-30 19:00:39.366876
6	1	72	2023-07-30 19:00:44.303444
7	1	90	2023-07-31 07:43:30.91134
8	2	100	2023-07-31 07:43:51.541617
9	3	100	2023-07-31 07:43:57.203818
10	3	110	2023-07-31 07:44:12.019117
11	1	62	2023-07-31 13:29:15.886166
\.


--
-- TOC entry 2911 (class 0 OID 167357)
-- Dependencies: 201
-- Data for Name: product; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.product (id, name, description, quantity, price_date) FROM stdin;
2	Апельсины		16	2023-07-31 07:43:51.541617
3	Огурцы		25	2023-07-31 07:44:12.019117
1	Молоко		24	2023-07-31 13:29:15.886166
\.


--
-- TOC entry 2914 (class 0 OID 167378)
-- Dependencies: 204
-- Data for Name: product_tag; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.product_tag (product_id, tag_id) FROM stdin;
1	1
1	2
2	1
2	4
3	1
3	6
\.


--
-- TOC entry 2913 (class 0 OID 167370)
-- Dependencies: 203
-- Data for Name: tag; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey.tag (id, name) FROM stdin;
2	Молочные продукты
4	Фрукты
1	Продукты
6	Овощи
\.


--
-- TOC entry 2909 (class 0 OID 167342)
-- Dependencies: 199
-- Data for Name: user; Type: TABLE DATA; Schema: grey; Owner: postgres
--

COPY grey."user" (id, username, firstname, lastname, fullname, age, is_married, password) FROM stdin;
1	noname	Иван	Иванов	Иван Иванов	21	f	75665a724b55444a7a65486743347143374b77656d4b706e34776e7558564c66f7c3bc1d808e04732adf679965ccc34ca7ae3441
2	admin	Петр	Петров	Петр Петров	27	t	75665a724b55444a7a65486743347143374b77656d4b706e34776e7558564c66f7c3bc1d808e04732adf679965ccc34ca7ae3441
\.


--
-- TOC entry 2907 (class 0 OID 167332)
-- Dependencies: 197
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
2	f
\.


--
-- TOC entry 2935 (class 0 OID 0)
-- Dependencies: 207
-- Name: cart_id_seq; Type: SEQUENCE SET; Schema: grey; Owner: postgres
--

SELECT pg_catalog.setval('grey.cart_id_seq', 16, true);


--
-- TOC entry 2936 (class 0 OID 0)
-- Dependencies: 210
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: grey; Owner: postgres
--

SELECT pg_catalog.setval('grey.order_id_seq', 13, true);


--
-- TOC entry 2937 (class 0 OID 0)
-- Dependencies: 205
-- Name: price_id_seq; Type: SEQUENCE SET; Schema: grey; Owner: postgres
--

SELECT pg_catalog.setval('grey.price_id_seq', 12, true);


--
-- TOC entry 2938 (class 0 OID 0)
-- Dependencies: 200
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: grey; Owner: postgres
--

SELECT pg_catalog.setval('grey.product_id_seq', 3, true);


--
-- TOC entry 2939 (class 0 OID 0)
-- Dependencies: 202
-- Name: tag_id_seq; Type: SEQUENCE SET; Schema: grey; Owner: postgres
--

SELECT pg_catalog.setval('grey.tag_id_seq', 6, true);


--
-- TOC entry 2940 (class 0 OID 0)
-- Dependencies: 198
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: grey; Owner: postgres
--

SELECT pg_catalog.setval('grey.user_id_seq', 9, true);


--
-- TOC entry 2770 (class 2606 OID 167430)
-- Name: cart_item cart_item_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart_item
    ADD CONSTRAINT cart_item_pkey PRIMARY KEY (product_id, cart_id, price_id);


--
-- TOC entry 2766 (class 2606 OID 167417)
-- Name: cart cart_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart
    ADD CONSTRAINT cart_pkey PRIMARY KEY (id);


--
-- TOC entry 2774 (class 2606 OID 167464)
-- Name: order_item order_item_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.order_item
    ADD CONSTRAINT order_item_pkey PRIMARY KEY (product_id, order_id, price_id);


--
-- TOC entry 2772 (class 2606 OID 167453)
-- Name: order order_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey."order"
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- TOC entry 2764 (class 2606 OID 167404)
-- Name: price price_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.price
    ADD CONSTRAINT price_pkey PRIMARY KEY (id);


--
-- TOC entry 2754 (class 2606 OID 167367)
-- Name: product product_name_key; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.product
    ADD CONSTRAINT product_name_key UNIQUE (name);


--
-- TOC entry 2756 (class 2606 OID 167365)
-- Name: product product_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- TOC entry 2762 (class 2606 OID 167382)
-- Name: product_tag product_tag_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.product_tag
    ADD CONSTRAINT product_tag_pkey PRIMARY KEY (product_id, tag_id);


--
-- TOC entry 2758 (class 2606 OID 167377)
-- Name: tag tag_name_key; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.tag
    ADD CONSTRAINT tag_name_key UNIQUE (name);


--
-- TOC entry 2760 (class 2606 OID 167375)
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- TOC entry 2768 (class 2606 OID 167419)
-- Name: cart user_id_unq; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart
    ADD CONSTRAINT user_id_unq UNIQUE (user_id);


--
-- TOC entry 2750 (class 2606 OID 167352)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2752 (class 2606 OID 167354)
-- Name: user user_username_key; Type: CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey."user"
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- TOC entry 2748 (class 2606 OID 167336)
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- TOC entry 2780 (class 2606 OID 167436)
-- Name: cart_item cart_item_cart_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart_item
    ADD CONSTRAINT cart_item_cart_id_fkey FOREIGN KEY (cart_id) REFERENCES grey.cart(id) ON DELETE CASCADE;


--
-- TOC entry 2781 (class 2606 OID 167441)
-- Name: cart_item cart_item_price_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart_item
    ADD CONSTRAINT cart_item_price_id_fkey FOREIGN KEY (price_id) REFERENCES grey.price(id) ON DELETE CASCADE;


--
-- TOC entry 2779 (class 2606 OID 167431)
-- Name: cart_item cart_item_product_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart_item
    ADD CONSTRAINT cart_item_product_id_fkey FOREIGN KEY (product_id) REFERENCES grey.product(id) ON DELETE CASCADE;


--
-- TOC entry 2778 (class 2606 OID 167420)
-- Name: cart cart_user_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.cart
    ADD CONSTRAINT cart_user_id_fkey FOREIGN KEY (user_id) REFERENCES grey."user"(id) ON DELETE CASCADE;


--
-- TOC entry 2784 (class 2606 OID 167470)
-- Name: order_item order_item_order_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.order_item
    ADD CONSTRAINT order_item_order_id_fkey FOREIGN KEY (order_id) REFERENCES grey."order"(id) ON DELETE CASCADE;


--
-- TOC entry 2785 (class 2606 OID 167475)
-- Name: order_item order_item_price_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.order_item
    ADD CONSTRAINT order_item_price_id_fkey FOREIGN KEY (price_id) REFERENCES grey.price(id) ON DELETE CASCADE;


--
-- TOC entry 2783 (class 2606 OID 167465)
-- Name: order_item order_item_product_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.order_item
    ADD CONSTRAINT order_item_product_id_fkey FOREIGN KEY (product_id) REFERENCES grey.product(id) ON DELETE CASCADE;


--
-- TOC entry 2782 (class 2606 OID 167454)
-- Name: order order_user_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey."order"
    ADD CONSTRAINT order_user_id_fkey FOREIGN KEY (user_id) REFERENCES grey."user"(id) ON DELETE CASCADE;


--
-- TOC entry 2777 (class 2606 OID 167405)
-- Name: price price_product_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.price
    ADD CONSTRAINT price_product_id_fkey FOREIGN KEY (product_id) REFERENCES grey.product(id) ON DELETE CASCADE;


--
-- TOC entry 2775 (class 2606 OID 167383)
-- Name: product_tag product_tag_product_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.product_tag
    ADD CONSTRAINT product_tag_product_id_fkey FOREIGN KEY (product_id) REFERENCES grey.product(id) ON DELETE CASCADE;


--
-- TOC entry 2776 (class 2606 OID 167388)
-- Name: product_tag product_tag_tag_id_fkey; Type: FK CONSTRAINT; Schema: grey; Owner: postgres
--

ALTER TABLE ONLY grey.product_tag
    ADD CONSTRAINT product_tag_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES grey.tag(id) ON DELETE CASCADE;


-- Completed on 2023-07-31 19:02:05

--
-- PostgreSQL database dump complete
--

