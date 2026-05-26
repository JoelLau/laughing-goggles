-- name: CreateLedgerEntry :one
INSERT INTO ledger_entries(event_id, account_id, amount_micro)
  VALUES (@event_id, @account_id, @amount_micro)
RETURNING
  *;

