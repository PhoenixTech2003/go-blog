-- +goose Up
CREATE TABLE articles (
    id UUID PRIMARY KEY,
    author TEXT NOT NULL,
    article_text TEXT NOT NULL,
    published_at TIMESTAMP
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE articles;
