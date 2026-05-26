-- name: CreateAccount :exec
INSERT INTO accounts (id)
  VALUES (@account_id)
RETURNING
  *;

-- name: GetAccountByID :one
SELECT
  acc.id AS account_id
  , COALESCE(SUM(ledger.amount_micro) , 0)::NUMERIC / 1000000 AS balance
FROM
  accounts acc
  LEFT JOIN ledger_entries ledger ON ledger.account_id = acc.id
WHERE
  acc.id = @account_id
GROUP BY
  acc.id;

