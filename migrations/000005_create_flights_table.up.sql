CREATE TABLE IF NOT EXISTS flights(
    id uuid,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    departure_id uuid,
    destination_id uuid,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_flights_airports_departure FOREIGN KEY (departure_id) REFERENCES airports (id) ON DELETE
    SET NULL,
        CONSTRAINT fk_flights_airports_destination FOREIGN KEY (destination_id) REFERENCES airports (id) ON DELETE
    SET NULL
);
