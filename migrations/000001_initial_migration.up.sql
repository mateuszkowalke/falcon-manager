CREATE TYPE role AS ENUM (
  'THERAPIST',
  'ADMIN',
  'USER'
);

CREATE TYPE sex AS ENUM (
  'MALE',
  'FEMALE',
  'OTHER'
);

CREATE TYPE meeting_status AS ENUM (
  'PENDING',
  'ARCHIVED',
  'CONFIRMED',
  'NOT_CONFIRMED',
  'ACTIVE'
);

CREATE TABLE app_user (
  id SERIAL PRIMARY KEY,
  email VARCHAR UNIQUE NOT NULL,
  username VARCHAR NOT NULL,
  first_name VARCHAR,
  last_name VARCHAR,
  avatar VARCHAR,
  sex sex,
  password VARCHAR NOT NULL,
  role role NOT NULL,
  blocked BOOLEAN NOT NULL,
  active BOOLEAN NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE therapist (
  id SERIAL PRIMARY KEY,
  user_id INT,
  work_start time NOT NULL,
  work_end time NOT NULL,
  experience INT NOT NULL,
  rating INT
);

ALTER TABLE therapist ADD FOREIGN KEY (user_id) REFERENCES app_user (id);

CREATE TABLE rating (
  id SERIAL PRIMARY KEY,
  therapist_id INT,
  score INT
);

ALTER TABLE rating ADD FOREIGN KEY (therapist_id) REFERENCES therapist (id);

CREATE TABLE meeting (
  id SERIAL PRIMARY KEY,
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL,
  user_id INT,
  therapist_id INT,
  status meeting_status NOT NULL,
  bill INT NOT NULL
);

ALTER TABLE meeting ADD FOREIGN KEY (user_id) REFERENCES app_user (id);
ALTER TABLE meeting ADD FOREIGN KEY (therapist_id) REFERENCES therapist (id);

CREATE TABLE specialization (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

CREATE TABLE problem (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

CREATE TABLE problem_specialization (
  id SERIAL PRIMARY KEY,
  problem_id INT,
  specialization_id INT
);

ALTER TABLE problem_specialization ADD FOREIGN KEY (problem_id) REFERENCES problem (id);
ALTER TABLE problem_specialization ADD FOREIGN KEY (specialization_id) REFERENCES specialization (id);

CREATE TABLE user_problem (
  id SERIAL PRIMARY KEY,
  user_id INT,
  problem_id INT
);

ALTER TABLE user_problem ADD FOREIGN KEY (user_id) REFERENCES app_user (id);
ALTER TABLE user_problem ADD FOREIGN KEY (problem_id) REFERENCES problem (id);

CREATE TABLE therapist_specialization (
  id SERIAL PRIMARY KEY,
  therapist_id INT,
  specialization_id INT
);

ALTER TABLE therapist_specialization ADD FOREIGN KEY (therapist_id) REFERENCES therapist (id);
ALTER TABLE therapist_specialization ADD FOREIGN KEY (specialization_id) REFERENCES specialization (id);
