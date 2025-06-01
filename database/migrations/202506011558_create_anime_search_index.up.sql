CREATE EXTENSION IF NOT EXISTS pg_trgm;

--bun:split

CREATE INDEX idx_anime_titles_search ON anime
USING GIN ((title || ' ' || COALESCE(title_english, '') || ' ' || title_japanese || ' ' || COALESCE(title_synonyms, '')) gin_trgm_ops);