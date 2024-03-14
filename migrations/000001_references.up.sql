create schema if not exists "references";

create table if not exists "references"."nationalities"
(
    "id" serial primary key,
    "title" text not null,
    "is_deleted" boolean default false,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create table if not exists "references"."countries"
(
    "id" serial primary key,
    "title" text not null,
    "link_to_flag_img" text not null,
    "is_deleted" boolean default false,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create table if not exists "references"."cities"
(
    "id" serial primary key,
    "country_id" integer references "references".countries(id) not null,
    "title" text not null,
    "link_to_flag_img" text not null,
    "is_deleted" boolean default false,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create table if not exists "references"."currency"
(
    "id" serial primary key,
    "title" text not null,
    "is_deleted" boolean default false,
    "create_date" timestamptz default timezone('UTC', current_timestamp),
    "update_date" timestamptz
);

create index if not exists idx_references_nationalities_title on "references".nationalities("title");

create index if not exists idx_references_countries_title on "references".countries("title");

create index if not exists idx_references_cities_title on "references".cities("title");
create index if not exists idx_references_cities_country_id on "references".cities("country_id");

create index if not exists idx_references_currency_title on "references".currency("title");