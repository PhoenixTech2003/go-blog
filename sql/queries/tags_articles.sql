-- name: AddTagToArticle :one
INSERT INTO tagsArticles (id ,aritcle_id, tag_id, created_at, updated_at) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    now(),
    now()
)

RETURNING *;