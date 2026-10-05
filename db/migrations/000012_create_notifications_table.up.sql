CREATE TYPE public.notification_type AS ENUM (
    'event_reminder',
    'registration',
    'community',
    'event_updated',
    'discussion_reply'
);

CREATE TABLE public.notifications (
    id integer NOT NULL,
    account_id integer NOT NULL,
    title character varying(200) NOT NULL,
    message text,
    type public.notification_type,
    read_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);