CREATE TYPE coordinates AS (latitude float, longitude float);
ALTER TABLE airports
ADD COLUMN location coordinates NOT NULL;
