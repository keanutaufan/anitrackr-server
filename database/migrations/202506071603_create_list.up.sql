CREATE TABLE lists
(
    anime_id   int         NOT NULL,
    user_id    int         NOT NULL,
    name       text        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (anime_id, user_id),
    FOREIGN KEY (anime_id) REFERENCES anime (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    CHECK (name IN ('watching', 'completed', 'on_hold', 'dropped', 'plan_to_watch'))
);

-- bun:split

CREATE INDEX IF NOT EXISTS idx_lists_name
ON lists
USING BTREE (name);