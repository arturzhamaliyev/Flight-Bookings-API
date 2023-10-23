CREATE TABLE IF NOT EXISTS tickets(
    id uuid,
    flight_id uuid,
    user_id uuid,
    rank_id uuid,
    price NUMERIC DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_tickets_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE
    SET NULL,
        CONSTRAINT fk_tickets_ranks FOREIGN KEY (rank_id) REFERENCES ranks (id) ON DELETE
    SET NULL
);
