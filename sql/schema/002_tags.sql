-- +goose Up
CREATE TABLE tags (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE tags;

