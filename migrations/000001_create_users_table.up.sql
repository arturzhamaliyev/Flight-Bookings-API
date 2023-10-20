CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4 (),
    first_name VARCHAR (255) NOT NULL,
    last_name VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL CHECK (password <> ''),
    email VARCHAR NOT NULL,
    country VARCHAR (255) NOT NULL CHECK (country <> ''),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT email_unique UNIQUE (email)
);
