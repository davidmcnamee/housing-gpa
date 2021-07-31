-- Enter migration here

create schema sh_public;
create schema sh_private;

create table sh_private.users (
  id serial primary key,
  name text not null,
  email text NOT NULL,
  password text NOT NULL,
  created_at timestamp default now(),
)

CREATE TABLE sh_public.building (
  id serial primary key,
  name text NOT NULL,
  address text NOT NULL,
  postcode text NOT NULL,
  created_at timestamp default now(),
)

CREATE TABLE sh_public.apartment (
  id serial primary key,
  building_id serial NOT NULL REFERENCES sh_public.building(id),
  number text NOT NULL,
  size integer NOT NULL,
  price integer NOT NULL,
  created_at timestamp default now(),
)
