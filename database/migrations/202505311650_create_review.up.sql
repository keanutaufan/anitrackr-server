CREATE TABLE reviews (
    id serial PRIMARY KEY,
    title text NOT NULL,
    body text NOT NULL,
    anime_id int NOT NULL,
    user_id int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    FOREIGN KEY (anime_id) REFERENCES anime (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    UNIQUE (anime_id, user_id)
);