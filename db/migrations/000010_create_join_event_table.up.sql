CREATE TABLE public.join_event (
    account_id integer NOT NULL,
    event_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.join_event
    ADD CONSTRAINT join_event_pkey PRIMARY KEY (account_id, event_id);