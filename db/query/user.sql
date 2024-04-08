
-- name: CreateUser :one
INSERT INTO "user" (hashed_password, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
where email = $1
LIMIT 1;