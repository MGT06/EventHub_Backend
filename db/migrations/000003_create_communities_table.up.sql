CREATE TABLE public.communities (
    id integer NOT NULL,
    community_name character varying(150) NOT NULL,
    description text,
    image_community_url character varying(500),
    is_active boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

ALTER TABLE ONLY public.communities
    ADD CONSTRAINT communities_pkey PRIMARY KEY (id);