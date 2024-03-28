-- name: GeUsers :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users ORDER BY UserId;

-- name: CreateUser :one
INSERT INTO users (
    Name, LastName, Email, UserName, Password,  Enabled
    ) VALUES (
        $1, $2, $3, $4, $5, $6
    )
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM products
WHERE UserId = $1;
