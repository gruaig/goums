-- name: GeUsers :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
    name, lastname, email, username, password,  enabled
    ) VALUES (
        $1, $2, $3, $4, $5, $6
    )
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

