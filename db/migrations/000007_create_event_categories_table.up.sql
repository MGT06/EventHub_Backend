CREATE TABLE public.event_categories (
    event_id integer NOT NULL,
    category_id integer NOT NULL
);

ALTER TABLE ONLY public.event_categories
    ADD CONSTRAINT event_categories_pkey PRIMARY KEY (event_id, category_id);