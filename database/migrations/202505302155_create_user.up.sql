CREATE TABLE users
(
    id         serial PRIMARY KEY,
    email      text        NOT NULL,
    uid        text        NOT NULL,
    name       text        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (email),
    UNIQUE (uid)
)