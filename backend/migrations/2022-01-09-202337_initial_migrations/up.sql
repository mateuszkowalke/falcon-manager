CREATE TYPE "sex" AS ENUM (
  'MALE',
  'FEMALE',
  'UNKNOWN'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar,
  "name" varchar,
  "password" varchar,
  "address_id" int,
  "photo_file" varchar,
  "role" int,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "users_breeding_projects" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "breeding_project_id" int
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "permissions" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "roles_permissions" (
  "id" SERIAL PRIMARY KEY,
  "role_id" int,
  "permission_id" int
);

CREATE TABLE "breeding_projects" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "vet_reg_no" varchar,
  "owner_id" int,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "falcons" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "ring" varchar,
  "species_id" int,
  "sex" sex,
  "birth_date" timestamp with time zone,
  "source" varchar,
  "aviary_id" int,
  "father" int,
  "mother" int,
  "width_young" int,
  "length_young" int,
  "weight_young" int,
  "width_old" int,
  "length_old" int,
  "weight_old" int,
  "notes" varchar,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone,
  "breeding_project_id" int
);

CREATE TABLE "pairs" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "male_id" int,
  "female_id" int,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "species" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "latin" varchar
);

CREATE TABLE "documents" (
  "id" SERIAL PRIMARY KEY,
  "falcon_id" int,
  "document_type_id" int,
  "document_number" varchar,
  "scan_file" varchar,
  "raw_file" varchar,
  "office_id" int,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "document_types" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "aviaries" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "capacity" varchar,
  "last_cleaned" timestamp with time zone,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone,
  "breeding_project_id" int
);

CREATE TABLE "photos" (
  "id" SERIAL PRIMARY KEY,
  "falcon_id" int,
  "file" varchar
);

CREATE TABLE "offices" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "office_type_id" int,
  "address_id" int,
  "breeding_project_id" int
);

CREATE TABLE "office_types" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "addresses" (
  "id" SERIAL PRIMARY KEY,
  "street" varchar,
  "no" varchar,
  "zip_code" varchar,
  "city" varchar,
  "country" varchar,
  "owner_id" int,
  "breeding_project_id" int
);

ALTER TABLE "users" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id");

ALTER TABLE "users_breeding_projects" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_breeding_projects" ADD FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects" ("id");

ALTER TABLE "roles_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "roles_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "breeding_projects" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "falcons" ADD FOREIGN KEY ("species_id") REFERENCES "species" ("id");

ALTER TABLE "falcons" ADD FOREIGN KEY ("aviary_id") REFERENCES "aviaries" ("id");

ALTER TABLE "falcons" ADD FOREIGN KEY ("father") REFERENCES "falcons" ("id");

ALTER TABLE "falcons" ADD FOREIGN KEY ("mother") REFERENCES "falcons" ("id");

ALTER TABLE "falcons" ADD FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects" ("id");

ALTER TABLE "pairs" ADD FOREIGN KEY ("male_id") REFERENCES "falcons" ("id");

ALTER TABLE "pairs" ADD FOREIGN KEY ("female_id") REFERENCES "falcons" ("id");

ALTER TABLE "documents" ADD FOREIGN KEY ("falcon_id") REFERENCES "falcons" ("id");

ALTER TABLE "documents" ADD FOREIGN KEY ("document_type_id") REFERENCES "document_types" ("id");

ALTER TABLE "documents" ADD FOREIGN KEY ("office_id") REFERENCES "offices" ("id");

ALTER TABLE "aviaries" ADD FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects" ("id");

ALTER TABLE "photos" ADD FOREIGN KEY ("falcon_id") REFERENCES "falcons" ("id");

ALTER TABLE "offices" ADD FOREIGN KEY ("office_type_id") REFERENCES "office_types" ("id");

ALTER TABLE "offices" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "offices" ADD FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects" ("id");
