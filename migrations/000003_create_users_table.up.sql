CREATE TABLE IF NOT EXISTS users(
    id uuid,
    phone VARCHAR (50),
    email VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT email_unique UNIQUE (email)
);
