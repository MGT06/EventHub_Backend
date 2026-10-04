
CREATE TABLE public.community_members (
    account_id integer NOT NULL,
    community_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.community_members
    ADD CONSTRAINT community_members_pkey PRIMARY KEY (account_id, community_id);