CREATE TABLE public.location_event (
    id integer NOT NULL,
    city character varying(150) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.location_event
    ADD CONSTRAINT location_event_pkey PRIMARY KEY (id);