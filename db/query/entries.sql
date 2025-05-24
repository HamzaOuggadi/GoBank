-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntriesByAccountId :many
SELECT * FROM entries
WHERE account_id=$1
ORDER BY id
LIMIT $1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1;