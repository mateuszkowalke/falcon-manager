ALTER TABLE "users" DROP CONSTRAINT "users_address_id_fkey";

ALTER TABLE "users" DROP CONSTRAINT "users_role_fkey";

ALTER TABLE "users_breeding_projects" DROP CONSTRAINT "users_breeding_projects_user_id_fkey";

ALTER TABLE "users_breeding_projects" DROP CONSTRAINT "users_breeding_projects_breeding_project_id_fkey";

ALTER TABLE "roles_permissions" DROP CONSTRAINT "roles_permissions_role_id_fkey";

ALTER TABLE "roles_permissions" DROP CONSTRAINT "roles_permissions_permission_id_fkey";

ALTER TABLE "breeding_projects" DROP CONSTRAINT "breeding_projects_owner_id_fkey";

ALTER TABLE "falcons" DROP CONSTRAINT "falcons_species_id_fkey";

ALTER TABLE "falcons" DROP CONSTRAINT "falcons_aviary_id_fkey";

ALTER TABLE "falcons" DROP CONSTRAINT "falcons_sex_id_fkey";

ALTER TABLE "falcons" DROP CONSTRAINT "falcons_father_fkey";

ALTER TABLE "falcons" DROP CONSTRAINT "falcons_mother_fkey";

ALTER TABLE "falcons" DROP CONSTRAINT "falcons_breeding_project_id_fkey";

ALTER TABLE "pairs" DROP CONSTRAINT "pairs_male_id_fkey";

ALTER TABLE "pairs" DROP CONSTRAINT "pairs_female_id_fkey";

ALTER TABLE "documents" DROP CONSTRAINT "documents_falcon_id_fkey";

ALTER TABLE "documents" DROP CONSTRAINT "documents_document_type_id_fkey";

ALTER TABLE "documents" DROP CONSTRAINT "documents_office_id_fkey";

ALTER TABLE "aviaries" DROP CONSTRAINT "aviaries_breeding_project_id_fkey";

ALTER TABLE "photos" DROP CONSTRAINT "photos_falcon_id_fkey";

ALTER TABLE "offices" DROP CONSTRAINT "offices_office_type_id_fkey";

ALTER TABLE "offices" DROP CONSTRAINT "offices_address_id_fkey";

ALTER TABLE "offices" DROP CONSTRAINT "offices_breeding_project_id_fkey";

ALTER TABLE "addresses" DROP CONSTRAINT "addresses_owner_id_fkey";

ALTER TABLE "addresses" DROP CONSTRAINT "addresses_breeding_project_id_fkey";

DROP TABLE IF EXISTS "addresses";

DROP TABLE IF EXISTS "offices";

DROP TABLE IF EXISTS "office_types";

DROP TABLE IF EXISTS "photos";

DROP TABLE IF EXISTS "aviaries";

DROP TABLE IF EXISTS "document_types";

DROP TABLE IF EXISTS "documents";

DROP TABLE IF EXISTS "species";

DROP TABLE IF EXISTS "pairs";

DROP TABLE IF EXISTS "falcons";

DROP TABLE IF EXISTS "breeding_projects";

DROP TABLE IF EXISTS "permissions";

DROP TABLE IF EXISTS "roles";

DROP TABLE IF EXISTS "roles_permissions";

DROP TABLE IF EXISTS "users_breeding_projects";

DROP TABLE IF EXISTS "users";

DROP TABLE IF EXISTS "sex";
