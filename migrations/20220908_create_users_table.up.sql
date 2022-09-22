CREATE TABLE users (
       id serial PRIMARY KEY,
       first_name varchar(100),
       last_name varchar(100),
       email varchar(255),
       hash_password varchar(255),
       created_at timestamptz NOT NULL DEFAULT now(),
       updated_at timestamptz
);