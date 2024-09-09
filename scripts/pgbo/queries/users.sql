-- name: CreateUser :one
INSERT INTO "users" (email, username, password, created_at, updated_at)
VALUES (@email, @username, @password, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING users.*;

-- name: UpdateUser :one
UPDATE "users"
SET email = @email, username = @username, password = @password, updated_at = CURRENT_TIMESTAMP
WHERE username = @username
RETURNING "GUID";

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE username = @username;

-- name: GetByUsername :one
SELECT "email", "username", "password"
FROM users
WHERE username = @username;