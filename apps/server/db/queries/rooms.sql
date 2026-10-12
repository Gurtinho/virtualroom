-- name: CreateRoom :one
INSERT INTO rooms (name, capacity)
VALUES ($1, $2)
RETURNING *;

-- name: GetRoom :one
SELECT * FROM rooms WHERE id = $1;

-- name: ListRooms :many
SELECT * FROM rooms ORDER BY id;

-- name: UpdateRoom :one
UPDATE rooms SET name = $2, capacity = $3
WHERE id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms WHERE id = $1;