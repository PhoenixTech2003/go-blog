-- +goose Up
CREATE TABLE tagsarticles (
    id UUID PRIMARY KEY,
    aritcle_id UUID NOT NULL REFERENCES articles(id),
    tag_id UUID NOT NULL REFERENCES tags(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE tagsarticles;