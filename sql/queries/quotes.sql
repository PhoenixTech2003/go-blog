-- name: CreateQoute :one
INSERT INTO quotes(id,author,qoute_text, created_at,updated_at)
VALUES(
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)

RETURNING *;

-- name: ListQuotes :many
SELECT * FROM quotes;
