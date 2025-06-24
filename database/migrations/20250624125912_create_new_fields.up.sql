ALTER TABLE reviews
ADD COLUMN is_liked BOOLEAN NOT NULL DEFAULT FALSE;

--bun:split

ALTER TABLE lists
ADD COLUMN episode_watched INT NOT NULL DEFAULT 0;

-- bun:split

ALTER TABLE ratings
ADD COLUMN episode_watched INT NOT NULL DEFAULT 0;