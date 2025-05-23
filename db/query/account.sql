-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListAccounts :many
SELECT * from accounts
ORDER BY id
LIMIT $1;

-- name: GetAccountById :one
SELECT * from accounts
WHERE id=$1 LIMIT 1;

-- name: UpdateAccountBalance :exec
UPDATE accounts
SET balance = $2
WHERE id = $1;

-- name: DeleteAccountById :exec
DELETE FROM accounts
WHERE id = $1;