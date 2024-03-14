drop index if exists "references".idx_geo_cities_country_id;
drop index if exists "references".idx_geo_cities_title;
drop index if exists "references".idx_geo_countries_title;
drop index if exists "references".idx_geo_nationalities_title;
drop index if exists "references".idx_references_currency_title;

drop table if exists "references"."cities";
drop table if exists "references"."countries";
drop table if exists "references"."nationalities";
drop table if exists "references"."currency";

drop schema if exists "references";