CREATE TABLE users (
       id int64 PRIMARY KEY,
       first_name varchar(100),
       last_name varchar(100),
       email varchar(255),
       password varchar(255),
       created_at timestamptz NOT NULL DEFAULT now(),
       updated_at timestampz,
);