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

-- name: GetUserForUpdate :one
SELECT *
FROM "user"
WHERE "username" = $1 LIMIT 1 FOR UPDATE;

-- name: ListUser :many
SELECT *
FROM "user"
ORDER BY "username";

-- name: UpdateUser :one
UPDATE "user" 
SET "username"=$1, "email"=$2, "full_name"=$3, "password"=$4
WHERE "username" = $5
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE "username" = $1;