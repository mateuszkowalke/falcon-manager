table! {
    addresses (id) {
        id -> Int4,
        street -> Nullable<Varchar>,
        no -> Nullable<Varchar>,
        zip_code -> Nullable<Varchar>,
        city -> Nullable<Varchar>,
        country -> Nullable<Varchar>,
        owner_id -> Nullable<Int4>,
        breeding_project_id -> Nullable<Int4>,
    }
}

table! {
    aviaries (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
        capacity -> Nullable<Varchar>,
        last_cleaned -> Nullable<Timestamptz>,
        created_at -> Nullable<Timestamptz>,
        updated_at -> Nullable<Timestamptz>,
        breeding_project_id -> Nullable<Int4>,
    }
}

table! {
    breeding_projects (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
        vet_reg_no -> Nullable<Varchar>,
        owner_id -> Nullable<Int4>,
        created_at -> Nullable<Timestamptz>,
        updated_at -> Nullable<Timestamptz>,
    }
}

table! {
    document_types (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
    }
}

table! {
    documents (id) {
        id -> Int4,
        falcon_id -> Nullable<Int4>,
        document_type_id -> Nullable<Int4>,
        document_number -> Nullable<Varchar>,
        scan_file -> Nullable<Varchar>,
        raw_file -> Nullable<Varchar>,
        office_id -> Nullable<Int4>,
        created_at -> Nullable<Timestamptz>,
        updated_at -> Nullable<Timestamptz>,
    }
}

table! {
    falcons (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
        ring -> Nullable<Varchar>,
        species_id -> Nullable<Int4>,
        sex -> Nullable<Sex>,
        birth_date -> Nullable<Timestamptz>,
        source -> Nullable<Varchar>,
        aviary_id -> Nullable<Int4>,
        father -> Nullable<Int4>,
        mother -> Nullable<Int4>,
        width_young -> Nullable<Int4>,
        length_young -> Nullable<Int4>,
        weight_young -> Nullable<Int4>,
        width_old -> Nullable<Int4>,
        length_old -> Nullable<Int4>,
        weight_old -> Nullable<Int4>,
        notes -> Nullable<Varchar>,
        created_at -> Nullable<Timestamptz>,
        updated_at -> Nullable<Timestamptz>,
        breeding_project_id -> Nullable<Int4>,
    }
}

table! {
    office_types (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
    }
}

table! {
    offices (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
        office_type_id -> Nullable<Int4>,
        address_id -> Nullable<Int4>,
        breeding_project_id -> Nullable<Int4>,
    }
}

table! {
    pairs (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
        male_id -> Nullable<Int4>,
        female_id -> Nullable<Int4>,
        created_at -> Nullable<Timestamptz>,
        updated_at -> Nullable<Timestamptz>,
    }
}

table! {
    permissions (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
    }
}

table! {
    photos (id) {
        id -> Int4,
        falcon_id -> Nullable<Int4>,
        file -> Nullable<Varchar>,
    }
}

table! {
    roles (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
    }
}

table! {
    roles_permissions (id) {
        id -> Int4,
        role_id -> Nullable<Int4>,
        permission_id -> Nullable<Int4>,
    }
}

table! {
    species (id) {
        id -> Int4,
        name -> Nullable<Varchar>,
        latin -> Nullable<Varchar>,
    }
}

table! {
    users (id) {
        id -> Int4,
        email -> Nullable<Varchar>,
        name -> Nullable<Varchar>,
        password -> Nullable<Varchar>,
        address_id -> Nullable<Int4>,
        photo_file -> Nullable<Varchar>,
        role -> Nullable<Int4>,
        created_at -> Nullable<Timestamptz>,
        updated_at -> Nullable<Timestamptz>,
    }
}

table! {
    users_breeding_projects (id) {
        id -> Int4,
        user_id -> Nullable<Int4>,
        breeding_project_id -> Nullable<Int4>,
    }
}

joinable!(addresses -> breeding_projects (breeding_project_id));
joinable!(aviaries -> breeding_projects (breeding_project_id));
joinable!(breeding_projects -> users (owner_id));
joinable!(documents -> document_types (document_type_id));
joinable!(documents -> falcons (falcon_id));
joinable!(documents -> offices (office_id));
joinable!(falcons -> aviaries (aviary_id));
joinable!(falcons -> breeding_projects (breeding_project_id));
joinable!(falcons -> species (species_id));
joinable!(offices -> addresses (address_id));
joinable!(offices -> breeding_projects (breeding_project_id));
joinable!(offices -> office_types (office_type_id));
joinable!(photos -> falcons (falcon_id));
joinable!(roles_permissions -> permissions (permission_id));
joinable!(roles_permissions -> roles (role_id));
joinable!(users -> roles (role));
joinable!(users_breeding_projects -> breeding_projects (breeding_project_id));
joinable!(users_breeding_projects -> users (user_id));

allow_tables_to_appear_in_same_query!(
    addresses,
    aviaries,
    breeding_projects,
    document_types,
    documents,
    falcons,
    office_types,
    offices,
    pairs,
    permissions,
    photos,
    roles,
    roles_permissions,
    species,
    users,
    users_breeding_projects,
);
