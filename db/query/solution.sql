
-- name: CreateSolution :one
INSERT INTO "solution" (condition, solution)
VALUES ($1, $2)
RETURNING *;

-- name: GetSolution :one
SELECT * FROM "solution"
where id = $1
LIMIT 1;