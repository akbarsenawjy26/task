-- name: CreateTask :one
INSERT INTO "task" (description, status, created_at, updated_at)
VALUES (@description, @status, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING task.*;

-- name: UpdateTask :one
UPDATE "task"
SET description = @description, status = @status, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING id;

-- name: DeleteTask :exec
DELETE FROM "task"
WHERE id = @id;

-- name: ListTask :many
SELECT id, description, status, created_at, updated_at
FROM task
ORDER BY created_at DESC;

-- name: ListTaskByStatus :many
SELECT id, description, status, created_at, updated_at
FROM task
WHERE status = @status
ORDER BY created_at DESC;