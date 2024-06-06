-- Version: 1.01
-- Description: Create table users
CREATE TABLE public.users (
    user_id uuid NOT NULL,
    "name" text NOT NULL,
    email text NOT NULL,
    roles text NOT NULL,
    password_hash text NOT NULL,
    guild text NULL,
    enabled bool NOT NULL,
    date_created timestamp NOT NULL,
    date_updated timestamp NOT NULL,

    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

-- Version: 1.02
-- Description: Create table galaxies
CREATE TABLE public.galaxies (
    galaxy_id uuid NOT NULL,
    galaxy_name text NOT NULL,
    owner_user_id uuid NOT NULL,

    CONSTRAINT galaxies_pk PRIMARY KEY (galaxy_id),
    CONSTRAINT galaxies_users_fk FOREIGN KEY (owner_user_id) REFERENCES public.users(user_id)
);

-- Version: 1.03
-- Description: Create table resources
CREATE TABLE public.resources (
    resource_id uuid NOT NULL,
    resource_name text NOT NULL,
    galaxy_id uuid NOT NULL,
    entered timestamp DEFAULT now() NOT NULL,
    entered_user_id uuid NOT NULL,
    resourcetype bpchar(63) NOT NULL,
    unavailable timestamp NULL,
    unavailable_user_id uuid NULL,
    verified timestamp NULL,
    verified_user_id uuid NULL,
    cr int2 DEFAULT 0 NOT NULL,
    cd int2 DEFAULT 0 NOT NULL,
    dr int2 DEFAULT 0 NOT NULL,
    fl int2 DEFAULT 0 NOT NULL,
    hr int2 DEFAULT 0 NOT NULL,
    ma int2 DEFAULT 0 NOT NULL,
    pe int2 DEFAULT 0 NOT NULL,
    oq int2 DEFAULT 0 NOT NULL,
    sr int2 DEFAULT 0 NOT NULL,
    ut int2 DEFAULT 0 NOT NULL,
    er int2 DEFAULT 0 NOT NULL,

    CONSTRAINT resources_pk PRIMARY KEY (resource_id),
    CONSTRAINT resources_galaxies_fk FOREIGN KEY (galaxy_id) REFERENCES public.galaxies(galaxy_id),
    CONSTRAINT resources_users_fk FOREIGN KEY (entered_user_id) REFERENCES public.users(user_id),
    CONSTRAINT resources_users_fk_1 FOREIGN KEY (unavailable_user_id) REFERENCES public.users(user_id),
    CONSTRAINT resources_users_fk_2 FOREIGN KEY (verified_user_id) REFERENCES public.users(user_id)
);
