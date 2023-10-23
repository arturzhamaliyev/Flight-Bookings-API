CREATE TABLE IF NOT EXISTS flights_prices(
    id uuid,
    flight_id uuid,
    rank_id uuid,
    price NUMERIC,
    PRIMARY KEY (id),
    CONSTRAINT fk_flights_prices_flight FOREIGN KEY (flight_id) REFERENCES flights (id) ON DELETE CASCADE,
    CONSTRAINT fk_flights_prices_rank FOREIGN KEY (rank_id) REFERENCES ranks (id) ON DELETE
    SET NULL
);
