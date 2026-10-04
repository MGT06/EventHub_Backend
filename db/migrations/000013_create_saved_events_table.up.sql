CREATE TABLE public.saved_events (
    account_id integer NOT NULL,
    event_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.saved_events
    ADD CONSTRAINT saved_events_pkey PRIMARY KEY (account_id, event_id);