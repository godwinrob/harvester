-- Version: 1.01
-- Description: Create table users
CREATE TABLE public.users (
    user_id uuid NOT NULL,
    "name" text NOT NULL,
    email text NOT NULL,
    roles text NOT NULL,
    password_hash text NOT NULL,
    guild text NULL,
    enabled bool default true NOT NULL,
    date_created timestamp DEFAULT now() NOT NULL,
    date_updated timestamp DEFAULT now() NOT NULL,

    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

-- Version: 1.02
-- Description: Create table galaxies
CREATE TABLE public.galaxies (
    galaxy_id uuid NOT NULL,
    galaxy_name text NOT NULL,
    owner_user_id uuid NOT NULL,
    enabled bool default true NOT NULL,
    date_created timestamp DEFAULT now() NOT NULL,
    date_updated timestamp DEFAULT now() NOT NULL,

    CONSTRAINT galaxies_pk PRIMARY KEY (galaxy_id)
);

-- Version: 1.03
-- Description: Create table resources
CREATE TABLE public.resources (
    resource_id uuid NOT NULL,
    resource_name text NOT NULL,
    galaxy_id uuid NOT NULL,
    added_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    added_user_id uuid NOT NULL,
    resource_type bpchar(63) NOT NULL,
    unavailable_at timestamp NULL,
    unavailable_user_id uuid NULL,
    verified bool DEFAULT false NOT NULL,
    verified_user_id uuid NULL,
    cr int2 DEFAULT 0 NOT NULL,
    cd int2 DEFAULT 0 NOT NULL,
    dr int2 DEFAULT 0 NOT NULL,
    fl int2 DEFAULT 0 NOT NULL,
    "hr" int2 DEFAULT 0 NOT NULL,
    ma int2 DEFAULT 0 NOT NULL,
    pe int2 DEFAULT 0 NOT NULL,
    oq int2 DEFAULT 0 NOT NULL,
    sr int2 DEFAULT 0 NOT NULL,
    ut int2 DEFAULT 0 NOT NULL,
    er int2 DEFAULT 0 NOT NULL,

    CONSTRAINT resources_pk PRIMARY KEY (resource_id)
);

-- Version: 1.04
-- Description: Create table resource_groups
CREATE TABLE public.resource_groups (
    resource_group  VARCHAR(63) NOT NULL,
    group_name      VARCHAR(255) NOT NULL,
    group_level     INT2 NOT NULL,
    group_order     INT2 NOT NULL DEFAULT 0,
    container_type  VARCHAR(63) NOT NULL DEFAULT '',

    CONSTRAINT resource_groups_pk PRIMARY KEY (resource_group)
);

-- Version: 1.05
-- Description: Create table resource_types
CREATE TABLE public.resource_types (
    resource_type       VARCHAR(63) NOT NULL,
    resource_type_name  VARCHAR(255) NOT NULL,
    resource_category   VARCHAR(63) NOT NULL DEFAULT '',
    resource_group      VARCHAR(63) NOT NULL DEFAULT '',
    enterable           BOOLEAN NOT NULL DEFAULT true,
    max_types           INT2 NOT NULL DEFAULT 1,
    cr_min INT2 DEFAULT 0 NOT NULL, cr_max INT2 DEFAULT 0 NOT NULL,
    cd_min INT2 DEFAULT 0 NOT NULL, cd_max INT2 DEFAULT 0 NOT NULL,
    dr_min INT2 DEFAULT 0 NOT NULL, dr_max INT2 DEFAULT 0 NOT NULL,
    fl_min INT2 DEFAULT 0 NOT NULL, fl_max INT2 DEFAULT 0 NOT NULL,
    hr_min INT2 DEFAULT 0 NOT NULL, hr_max INT2 DEFAULT 0 NOT NULL,
    ma_min INT2 DEFAULT 0 NOT NULL, ma_max INT2 DEFAULT 0 NOT NULL,
    pe_min INT2 DEFAULT 0 NOT NULL, pe_max INT2 DEFAULT 0 NOT NULL,
    oq_min INT2 DEFAULT 0 NOT NULL, oq_max INT2 DEFAULT 0 NOT NULL,
    sr_min INT2 DEFAULT 0 NOT NULL, sr_max INT2 DEFAULT 0 NOT NULL,
    ut_min INT2 DEFAULT 0 NOT NULL, ut_max INT2 DEFAULT 0 NOT NULL,
    er_min INT2 DEFAULT 0 NOT NULL, er_max INT2 DEFAULT 0 NOT NULL,
    container_type      VARCHAR(63) NOT NULL DEFAULT '',
    inventory_type      VARCHAR(63) NOT NULL DEFAULT '',
    specific_planet     INT2 NOT NULL DEFAULT 0,

    CONSTRAINT resource_types_pk PRIMARY KEY (resource_type)
);

-- Version: 1.06
-- Description: Create table resource_type_groups (maps types to all ancestor groups)
CREATE TABLE public.resource_type_groups (
    resource_type   VARCHAR(63) NOT NULL,
    resource_group  VARCHAR(63) NOT NULL,

    CONSTRAINT resource_type_groups_pk PRIMARY KEY (resource_type, resource_group)
);

-- Version: 1.07
-- Description: Add FK from resources.resource_type to resource_types
ALTER TABLE public.resources
    ADD CONSTRAINT resources_resource_type_fk
    FOREIGN KEY (resource_type) REFERENCES public.resource_types(resource_type);
