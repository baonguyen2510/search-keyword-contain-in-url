CREATE SCHEMA IF NOT exists "search";

CREATE TABLE IF NOT EXISTS search.keyword_rank ( 
	id TEXT NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	keyword TEXT,
	rank BIGINT DEFAULT 0,
	title text,
	url text,
	description text,
	PRIMARY KEY(id)
);