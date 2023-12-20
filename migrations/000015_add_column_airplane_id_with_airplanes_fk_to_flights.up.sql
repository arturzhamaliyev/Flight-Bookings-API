ALTER TABLE flights
ADD COLUMN airplane_id uuid,
    ADD CONSTRAINT airplane_id_fk FOREIGN KEY (airplane_id) REFERENCES airplanes (id) ON DELETE
SET NULL;
