-- name: CreateTag :one
INSERT INTO tags (id, name, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    now(),
    now()
)

RETURNING *;