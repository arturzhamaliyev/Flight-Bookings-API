CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS citext;
CREATE DOMAIN email AS citext CHECK (
    value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$'
);
CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4 (),
    first_name VARCHAR (255) NOT NULL,
    last_name VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL CHECK (password <> ''),
    email email NOT NULL,
    country VARCHAR (255) NOT NULL CHECK (country <> ''),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT id_unique UNIQUE (id),
    CONSTRAINT email_unique UNIQUE (email)
);
CREATE INDEX idx_users_first_name ON users (first_name);
CREATE INDEX idx_users_last_name ON users (last_name);
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_country ON users (country);