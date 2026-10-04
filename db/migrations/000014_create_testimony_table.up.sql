CREATE TABLE public.testimony (
    id integer NOT NULL,
    account_id integer NOT NULL,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.testimony
    ADD CONSTRAINT testimony_pkey PRIMARY KEY (id);