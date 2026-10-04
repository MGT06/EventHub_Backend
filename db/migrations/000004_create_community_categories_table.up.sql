CREATE TABLE public.community_categories (
    community_id integer NOT NULL,
    category_id integer NOT NULL
);

ALTER TABLE ONLY public.community_categories
    ADD CONSTRAINT community_categories_pkey PRIMARY KEY (community_id, category_id);