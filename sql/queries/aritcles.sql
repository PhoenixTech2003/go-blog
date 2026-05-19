-- name: CreateAritcle :one
INSERT INTO articles(id,author,article_text, created_at,updated_at, published_at)
VALUES(
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW(),
    NOW()
)

RETURNING *;

-- name: ListArticles :many
SELECT * FROM articles;
