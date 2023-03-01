-- name: CreateTodo :one
INSERT INTO todos (users_id, task, done)
VALUES ( $1, $2, $3) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodo :many
SELECT * FROM todos
WHERE users_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTodo :one
UPDATE todos
SET task = $2,
    done = $3
WHERE id = $1
RETURNING *;
