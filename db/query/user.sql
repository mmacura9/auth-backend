-- name: CreateUser :one
INSERT INTO "user" ("username", "email", "full_name", "password", "birth_date")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByUsername :one
SELECT *
FROM "user"
WHERE "username" = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM "user"
WHERE "email" = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE "username" = $1;