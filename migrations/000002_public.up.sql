create schema if not exists "public";

create table if not exists "public"."users"
(
    "id" serial primary key,
    "nationality_id" integer references "references".nationalities("id") not null,
    "full_name" text not null,
    "email" text not null,
    "phone_number" text,
    "gender" boolean not null,
    "is_deleted" boolean default false,
    "date_of_birth" time not null,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create table if not exists "public"."users_hash"
(
    "id" serial primary key,
    "user_id" integer references public.users("id") not null,
    "hash" text not null,
    "salt" text not null,
    "update_date" timestamptz
);


create table if not exists "public"."user_roles"
(
    "id" serial primary key,
    "user_id" integer references "public"."users"("id"),
    "role_id" integer references "references"."roles"("id"),
    "is_deleted" boolean default false,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create table if not exists "public"."search_settings"
(
    "id" serial primary key,
    "user_id" integer references "public"."users"("id"),
    "nationality_id" integer references "references"."nationalities"("id"),
    "country_id" integer references "references"."countries"("id"),
    "city_id" integer references "references"."cities"("id"),
    "currency_id" integer references "references".currency("id"),
    "min_sum" integer,
    "max_sum" integer,
    "gender" boolean,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create index if not exists idx_public_user_roles_user_id on "public"."user_roles"("user_id");
create index if not exists idx_public_user_roles_role_id on "public"."user_roles"("role_id");

create index if not exists idx_public_users_nationality_id on "public"."users"("nationality_id");
create index if not exists idx_public_users_gender on "public"."users"("gender");

create index if not exists idx_public_users_hash_user_id on "public"."users_hash"("user_id");

create index if not exists idx_public_search_settings_user_id on "public"."search_settings"("user_id");
create index if not exists idx_public_search_settings_nationality_id on "public"."search_settings"("nationality_id");
create index if not exists idx_public_search_settings_country_id on "public"."search_settings"("country_id");
create index if not exists idx_public_search_settings_city_id on "public"."search_settings"("city_id");
create index if not exists idx_public_search_settings_currency_id on "public"."search_settings"("currency_id");
create index if not exists idx_public_search_settings_gender on "public"."search_settings"("gender");