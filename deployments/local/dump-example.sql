--
-- PostgreSQL database dump
--

-- Dumped from database version 15.5 (Debian 15.5-1.pgdg120+1)
-- Dumped by pg_dump version 15.3

-- Started on 2024-01-23 02:21:36

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
-- TOC entry 217 (class 1259 OID 16414)
-- Name: placements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.placements (
    stock_id uuid NOT NULL,
    product_id uuid NOT NULL,
    count integer NOT NULL,
    reserved integer NOT NULL
);


ALTER TABLE public.placements OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16405)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    x integer NOT NULL,
    y integer NOT NULL,
    z integer NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16384)
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16396)
-- Name: stocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stocks (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    is_available boolean NOT NULL
);


ALTER TABLE public.stocks OWNER TO postgres;

--
-- TOC entry 3365 (class 0 OID 16414)
-- Dependencies: 217
-- Data for Name: placements; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.placements (stock_id, product_id, count, reserved) FROM stdin;
aa94d884-c976-4b79-9d7a-119d5e02c68c	5125470f-b113-4750-bdab-f2b6e6aa81fd	12	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	436cd4df-7012-49ce-af17-a066c136830a	11	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	79a6d760-6913-4bb6-8f5c-1afb51e96823	11	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	c3449ae9-748d-497f-9a90-c18d6d5eff06	5	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	2185749e-1005-46bb-8208-70796eb093a7	6	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	db124480-19f4-4d94-8669-4b0d426ab388	10	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	436cd4df-7012-49ce-af17-a066c136830a	13	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	436cd4df-7012-49ce-af17-a066c136830a	11	0
07d163a9-41cb-455e-8727-e82882b71c1b	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	12	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	436cd4df-7012-49ce-af17-a066c136830a	12	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	5125470f-b113-4750-bdab-f2b6e6aa81fd	5	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	8	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	10	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	7	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	8	0
07d163a9-41cb-455e-8727-e82882b71c1b	2185749e-1005-46bb-8208-70796eb093a7	6	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	2185749e-1005-46bb-8208-70796eb093a7	6	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	5125470f-b113-4750-bdab-f2b6e6aa81fd	8	0
07d163a9-41cb-455e-8727-e82882b71c1b	db124480-19f4-4d94-8669-4b0d426ab388	14	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	5125470f-b113-4750-bdab-f2b6e6aa81fd	6	0
07d163a9-41cb-455e-8727-e82882b71c1b	c3449ae9-748d-497f-9a90-c18d6d5eff06	11	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	5	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	db124480-19f4-4d94-8669-4b0d426ab388	9	0
07d163a9-41cb-455e-8727-e82882b71c1b	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	11	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	79a6d760-6913-4bb6-8f5c-1afb51e96823	12	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	11	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	db124480-19f4-4d94-8669-4b0d426ab388	7	0
07d163a9-41cb-455e-8727-e82882b71c1b	5125470f-b113-4750-bdab-f2b6e6aa81fd	13	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	c3449ae9-748d-497f-9a90-c18d6d5eff06	6	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	9	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	12	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	2185749e-1005-46bb-8208-70796eb093a7	12	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	db124480-19f4-4d94-8669-4b0d426ab388	12	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	5125470f-b113-4750-bdab-f2b6e6aa81fd	9	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	db124480-19f4-4d94-8669-4b0d426ab388	9	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	79a6d760-6913-4bb6-8f5c-1afb51e96823	9	0
07d163a9-41cb-455e-8727-e82882b71c1b	79a6d760-6913-4bb6-8f5c-1afb51e96823	12	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	7	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	12	0
07d163a9-41cb-455e-8727-e82882b71c1b	436cd4df-7012-49ce-af17-a066c136830a	6	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	12	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	12	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	13	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	6	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	2185749e-1005-46bb-8208-70796eb093a7	13	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	79a6d760-6913-4bb6-8f5c-1afb51e96823	6	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	2185749e-1005-46bb-8208-70796eb093a7	5	0
aa94d884-c976-4b79-9d7a-119d5e02c68c	79a6d760-6913-4bb6-8f5c-1afb51e96823	5	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	2185749e-1005-46bb-8208-70796eb093a7	10	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	db124480-19f4-4d94-8669-4b0d426ab388	5	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	8	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	436cd4df-7012-49ce-af17-a066c136830a	13	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	6794e5e1-b4ed-4f40-8e9f-d6eac226104a	14	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	c3449ae9-748d-497f-9a90-c18d6d5eff06	11	0
68f14a41-ea94-4f0a-b863-adeb9be846f7	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	7	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	79a6d760-6913-4bb6-8f5c-1afb51e96823	14	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	5125470f-b113-4750-bdab-f2b6e6aa81fd	12	0
620680a4-a672-42d8-9ca8-f802d0ca1ba1	c3449ae9-748d-497f-9a90-c18d6d5eff06	10	0
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	7	0
a63a42c4-9797-42b4-94c7-f42e208c6d91	436cd4df-7012-49ce-af17-a066c136830a	6	0
07d163a9-41cb-455e-8727-e82882b71c1b	8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	7	0
f43a76f7-7ef1-4d0e-8526-570f32670ef9	c3449ae9-748d-497f-9a90-c18d6d5eff06	9	0
\.


--
-- TOC entry 3364 (class 0 OID 16405)
-- Dependencies: 216
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, name, x, y, z) FROM stdin;
5125470f-b113-4750-bdab-f2b6e6aa81fd	Кросовки	1	1	1
8a5bb6e6-a717-472f-adff-dbfe40fdfbe4	Туфли	1	1	1
db124480-19f4-4d94-8669-4b0d426ab388	Пиджак	1	1	1
79a6d760-6913-4bb6-8f5c-1afb51e96823	Рубашка	1	1	1
6794e5e1-b4ed-4f40-8e9f-d6eac226104a	Свитер	1	1	1
c3449ae9-748d-497f-9a90-c18d6d5eff06	Шапка	1	1	1
1b4dcd9e-3a14-4bd6-b6dd-0764a5ce5c83	Кепка	1	1	1
436cd4df-7012-49ce-af17-a066c136830a	Толстовка	1	1	1
2185749e-1005-46bb-8208-70796eb093a7	Худи	1	1	1
\.


--
-- TOC entry 3362 (class 0 OID 16384)
-- Dependencies: 214
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
4	f
\.


--
-- TOC entry 3363 (class 0 OID 16396)
-- Dependencies: 215
-- Data for Name: stocks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.stocks (id, name, is_available) FROM stdin;
aa94d884-c976-4b79-9d7a-119d5e02c68c	Таганрог	t
f43a76f7-7ef1-4d0e-8526-570f32670ef9	Москва	t
cb89c91f-0c48-4ec7-a3de-cb7433991fa2	Псков	t
a63a42c4-9797-42b4-94c7-f42e208c6d91	Тверь	t
07d163a9-41cb-455e-8727-e82882b71c1b	Нижний Новгород	t
620680a4-a672-42d8-9ca8-f802d0ca1ba1	Новороссийск	t
68f14a41-ea94-4f0a-b863-adeb9be846f7	Коломна	t
\.


--
-- TOC entry 3217 (class 2606 OID 16418)
-- Name: placements placements_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.placements
    ADD CONSTRAINT placements_pkey PRIMARY KEY (stock_id, product_id);


--
-- TOC entry 3215 (class 2606 OID 16409)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 3211 (class 2606 OID 16388)
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- TOC entry 3213 (class 2606 OID 16400)
-- Name: stocks stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_pkey PRIMARY KEY (id);


--
-- TOC entry 3218 (class 2606 OID 16424)
-- Name: placements fk_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.placements
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- TOC entry 3219 (class 2606 OID 16419)
-- Name: placements fk_stock; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.placements
    ADD CONSTRAINT fk_stock FOREIGN KEY (stock_id) REFERENCES public.stocks(id);


-- Completed on 2024-01-23 02:21:36

--
-- PostgreSQL database dump complete
--

