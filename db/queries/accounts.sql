-- name: GetAccountByID :one
SELECT id
  , 0 AS balance
FROM
  accounts
WHERE
    id = sqlc.arg(account_id)
;