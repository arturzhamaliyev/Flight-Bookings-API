CREATE TABLE IF NOT EXISTS airports(
    id uuid,
    name VARCHAR (255) NOT NULL,
    city VARCHAR (255) NOT NULL,
    country VARCHAR (255) NOT NULL,
    PRIMARY KEY (id)
);
