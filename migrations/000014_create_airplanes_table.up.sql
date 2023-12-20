CREATE TABLE IF NOT EXISTS airplanes(
    id uuid,
    model VARCHAR(255) NOT NULL,
    number_of_seats INTEGER NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT model_unique UNIQUE (model)
);
