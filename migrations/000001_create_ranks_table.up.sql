CREATE TABLE IF NOT EXISTS ranks(
    id uuid,
    name VARCHAR (255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT name_unique UNIQUE (name)
);
