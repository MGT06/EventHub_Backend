CREATE TABLE public.event_discussions (
    id integer NOT NULL,
    event_id integer NOT NULL,
    account_id integer NOT NULL,
    parent_id integer,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

ALTER TABLE ONLY public.event_discussions
    ADD CONSTRAINT event_discussions_pkey PRIMARY KEY (id);