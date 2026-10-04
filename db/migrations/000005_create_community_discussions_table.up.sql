CREATE TABLE public.community_discussions (
    id integer NOT NULL,
    community_id integer NOT NULL,
    account_id integer NOT NULL,
    parent_id integer,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

ALTER TABLE ONLY public.community_discussions
    ADD CONSTRAINT community_discussions_pkey PRIMARY KEY (id);