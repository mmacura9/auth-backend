-- name: CreateSession :one
INSERT INTO "sessions" (
    "id",
    "username",
    "refresh_token",
    "user_agent",
    "client_ip",
    "is_blocked",
    "expires_at"
) VALUES (
$1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateSession :one
UPDATE "sessions" 
SET "refresh_token"=$1, "expires_at"=$2
WHERE "id" = $3
RETURNING *;

-- name: GetSessionByUsername :many
SELECT *
FROM "sessions"
WHERE "username" = $1;

-- name: GetSessionByID :one
SELECT *
FROM "sessions"
WHERE "id" = $1;