CREATE TABLE ratings (
    anime_id int NOT NULL,
    user_id int NOT NULL,
    score int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (anime_id, user_id),
    FOREIGN KEY (anime_id) REFERENCES anime (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    CHECK (score BETWEEN 1 AND 10)
);

--bun:split

CREATE INDEX IF NOT EXISTS idx_ratings_score
ON ratings
USING BTREE (score);