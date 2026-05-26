-- name: CreateEvent :one
INSERT INTO events(type, data)
    VALUES (@type, @data)
RETURNING
    *;

