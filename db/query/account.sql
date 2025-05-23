-- name CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name GetAccounts :one
SELECT * from accounts
WHERE id=$1 LIMIT 1;