create table countries
(
    "id" serial primary key,
    "name" varchar(50) not null,
    "flag_img_url" varchar(255) not null
);

create table cities
(
    "id" serial primary key,
    "country_id" integer references countries(id) not null,
    "name" varchar(150) not null,
    "flag_img_url" varchar(255) not null
);

create table gender
(
    "id" serial primary key,
    "name" varchar(50) not null,
);

create table currency
(
    "id" serial primary key,
    "name" varchar(50) not null
);

create table budgets
(
    "id" serial primary key,
    "currency_id" integer references currency(id) not null,
    "max_sum" integer,
    "min_sum" integer
);

create table users
(
    "id"                     serial primary key,
    "nationality_country_id" integer references countries (id) not null,
    "citizenship_country_id" integer references countries (id) not null,
    "gender_id"              integer references gender (id)    not null,
    "budget_id"              integer references budgets (id)   not null,
    "city_search_id"         integer references cities (id),
    "country_search_id"      integer references countries (id),
    "full_name"              varchar(100)                      not null,
    "phone_number"           varchar(100)                      not null,
    "email"                  varchar(255)                      not null,
    "password_hash"          varchar(255)                      not null
);
