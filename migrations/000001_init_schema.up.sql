CREATE TYPE "Sex" AS ENUM ('MALE', 'FEMALE', 'UNKNOWN');


CREATE TABLE "addresses" (
    "id" SERIAL NOT NULL,
    "street" TEXT NOT NULL,
    "no" TEXT NOT NULL,
    "zip_code" TEXT NOT NULL,
    "city" TEXT NOT NULL,
    "country" TEXT NOT NULL,
    "owner_id" INTEGER NOT NULL,
    "breeding_project_id" INTEGER,

    CONSTRAINT "addresses_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "aviaries" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "capacity" INTEGER NOT NULL,
    "last_cleaned" TIMESTAMP(3) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "breeding_project_id" INTEGER NOT NULL,

    CONSTRAINT "aviaries_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "birth_certs" (
    "id" SERIAL NOT NULL,
    "scan_file" TEXT NOT NULL,
    "raw_file" TEXT NOT NULL,
    "office_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "birth_certs_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "registrations" (
    "id" SERIAL NOT NULL,
    "scan_file" TEXT NOT NULL,
    "raw_file" TEXT NOT NULL,
    "office_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "registrations_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "cites" (
    "id" SERIAL NOT NULL,
    "scan_file" TEXT NOT NULL,
    "raw_file" TEXT NOT NULL,
    "office_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "cites_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "breeding_projects" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "vet_reg_no" TEXT NOT NULL,
    "owner_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "breeding_projects_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "documents" (
    "id" SERIAL NOT NULL,
    "falcon_id" INTEGER NOT NULL,
    "birth_cert_id" INTEGER NOT NULL,
    "cites_id" INTEGER NOT NULL,
    "registration_id" INTEGER NOT NULL,

    CONSTRAINT "documents_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "falcons" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "ring" TEXT NOT NULL,
    "species_id" INTEGER NOT NULL,
    "sex" "Sex" NOT NULL DEFAULT E'UNKNOWN',
    "birth_date" TIMESTAMP(3) NOT NULL,
    "source" TEXT NOT NULL,
    "aviary_id" INTEGER NOT NULL,
    "father_id" INTEGER,
    "mother_id" INTEGER,
    "width_young" INTEGER,
    "length_young" INTEGER,
    "weight_young" INTEGER,
    "width_old" INTEGER,
    "length_old" INTEGER,
    "weight_old" INTEGER,
    "notes" TEXT,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "breeding_project_id" INTEGER NOT NULL,

    CONSTRAINT "falcons_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "office_types" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "office_types_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "offices" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "office_type_id" INTEGER NOT NULL,
    "address_id" INTEGER NOT NULL,
    "breeding_project_id" INTEGER NOT NULL,

    CONSTRAINT "offices_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "pairs" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "male_id" INTEGER NOT NULL,
    "female_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "pairs_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "permissions" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "permissions_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "roles" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "roles_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "roles_permissions" (
    "role_id" INTEGER NOT NULL,
    "permission_id" INTEGER NOT NULL,

    CONSTRAINT "roles_permissions_pkey" PRIMARY KEY ("role_id","permission_id")
);


CREATE TABLE "photos" (
    "id" SERIAL NOT NULL,
    "falcon_id" INTEGER NOT NULL,
    "file" TEXT NOT NULL,

    CONSTRAINT "photos_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "species" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "latin" TEXT NOT NULL,

    CONSTRAINT "species_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "photo_file" TEXT,
    "role_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);


CREATE TABLE "users_breeding_projects" (
    "id" SERIAL NOT NULL,
    "user_id" INTEGER NOT NULL,
    "breeding_project_id" INTEGER NOT NULL,

    CONSTRAINT "users_breeding_projects_pkey" PRIMARY KEY ("id")
);


CREATE UNIQUE INDEX "permissions_name_key" ON "permissions"("name");


CREATE UNIQUE INDEX "roles_name_key" ON "roles"("name");


CREATE UNIQUE INDEX "users_name_key" ON "users"("name");


ALTER TABLE "addresses" ADD CONSTRAINT "addresses_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;


ALTER TABLE "addresses" ADD CONSTRAINT "addresses_breeding_project_id_fkey" FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "aviaries" ADD CONSTRAINT "aviaries_breeding_project_id_fkey" FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "birth_certs" ADD CONSTRAINT "birth_certs_office_id_fkey" FOREIGN KEY ("office_id") REFERENCES "offices"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "registrations" ADD CONSTRAINT "registrations_office_id_fkey" FOREIGN KEY ("office_id") REFERENCES "offices"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "cites" ADD CONSTRAINT "cites_office_id_fkey" FOREIGN KEY ("office_id") REFERENCES "offices"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "breeding_projects" ADD CONSTRAINT "breeding_projects_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "users"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "documents" ADD CONSTRAINT "documents_birth_cert_id_fkey" FOREIGN KEY ("birth_cert_id") REFERENCES "birth_certs"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "documents" ADD CONSTRAINT "documents_cites_id_fkey" FOREIGN KEY ("cites_id") REFERENCES "cites"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "documents" ADD CONSTRAINT "documents_falcon_id_fkey" FOREIGN KEY ("falcon_id") REFERENCES "falcons"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "documents" ADD CONSTRAINT "documents_registration_id_fkey" FOREIGN KEY ("registration_id") REFERENCES "registrations"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "falcons" ADD CONSTRAINT "falcons_aviary_id_fkey" FOREIGN KEY ("aviary_id") REFERENCES "aviaries"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "falcons" ADD CONSTRAINT "falcons_breeding_project_id_fkey" FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "falcons" ADD CONSTRAINT "falcons_father_fkey" FOREIGN KEY ("father_id") REFERENCES "falcons"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "falcons" ADD CONSTRAINT "falcons_mother_fkey" FOREIGN KEY ("mother_id") REFERENCES "falcons"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "falcons" ADD CONSTRAINT "falcons_species_id_fkey" FOREIGN KEY ("species_id") REFERENCES "species"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "offices" ADD CONSTRAINT "offices_address_id_fkey" FOREIGN KEY ("address_id") REFERENCES "addresses"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "offices" ADD CONSTRAINT "offices_breeding_project_id_fkey" FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "offices" ADD CONSTRAINT "offices_office_type_id_fkey" FOREIGN KEY ("office_type_id") REFERENCES "office_types"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "pairs" ADD CONSTRAINT "pairs_female_id_fkey" FOREIGN KEY ("female_id") REFERENCES "falcons"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "pairs" ADD CONSTRAINT "pairs_male_id_fkey" FOREIGN KEY ("male_id") REFERENCES "falcons"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "roles_permissions" ADD CONSTRAINT "roles_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "permissions"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "roles_permissions" ADD CONSTRAINT "roles_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "photos" ADD CONSTRAINT "photos_falcon_id_fkey" FOREIGN KEY ("falcon_id") REFERENCES "falcons"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "users" ADD CONSTRAINT "users_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "users_breeding_projects" ADD CONSTRAINT "users_breeding_projects_breeding_project_id_fkey" FOREIGN KEY ("breeding_project_id") REFERENCES "breeding_projects"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


ALTER TABLE "users_breeding_projects" ADD CONSTRAINT "users_breeding_projects_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

