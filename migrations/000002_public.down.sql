drop index if exists "public".idx_public_roles_role;
drop index if exists "public".idx_public_user_roles_user_id;
drop index if exists "public".idx_public_user_roles_role_id;
drop index if exists "public".idx_public_search_settings_gender;
drop index if exists "public".idx_public_search_settings_currency_id;
drop index if exists "public".idx_public_search_settings_city_id;
drop index if exists "public".idx_public_search_settings_country_id;
drop index if exists "public".idx_public_search_settings_nationality_id;
drop index if exists "public".idx_public_search_settings_user_id;
drop index if exists "public".idx_public_users_hash_user_id;
drop index if exists "public".idx_public_users_gender;
drop index if exists "public".idx_public_users_nationality_id;

drop table if exists "public"."user_roles";
drop table if exists "public"."roles";
drop table if exists "public"."users";
drop table if exists "public"."users_hash";
drop table if exists "public"."search_settings";

drop schema if exists "public";