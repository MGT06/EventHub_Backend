CREATE TABLE public.events (
    id integer NOT NULL,
    organizer_id integer NOT NULL,
    community_id integer,
    location_event_id integer NOT NULL,
    title character varying(200) NOT NULL,
    description text NOT NULL,
    image_event_url character varying(500) NOT NULL,
    start_at timestamp with time zone NOT NULL,
    end_at timestamp with time zone NOT NULL,
    format public.event_format DEFAULT 'offline'::public.event_format NOT NULL,
    capacity integer NOT NULL,
    speakers jsonb DEFAULT '[]'::jsonb NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    CONSTRAINT chk_events_capacity_positive CHECK ((capacity >= 0)),
    CONSTRAINT chk_events_end_after_start CHECK ((end_at > start_at))
);

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);