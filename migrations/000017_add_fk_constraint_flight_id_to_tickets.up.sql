ALTER TABLE tickets
ADD CONSTRAINT fk_tickets_flight FOREIGN KEY (flight_id) REFERENCES flights (id) ON DELETE CASCADE;
