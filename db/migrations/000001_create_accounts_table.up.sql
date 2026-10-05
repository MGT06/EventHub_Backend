CREATE TYPE public.role_type AS ENUM (
    'attendee',
    'organizer',
    'admin'
);

CREATE TYPE public.account_status AS ENUM (
    'active',
    'inactive',
    'suspended'
);


CREATE TABLE public.accounts (
    id integer NOT NULL PRIMARY KEY,
    name character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    bio text,
    user_location character varying(150),
    "position" character varying(150),
    password character varying(255) NOT NULL,
    role public.role_type DEFAULT 'attendee'::public.role_type NOT NULL,
    avatar_url character varying(255),
    status public.account_status DEFAULT 'active'::public.account_status,
    terms boolean,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_email_key UNIQUE (email);

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);