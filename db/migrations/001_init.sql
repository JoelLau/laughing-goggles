-- +goose Up
CREATE TABLE IF NOT EXISTS accounts (
  id BIGINT PRIMARY KEY
);

COMMENT ON TABLE accounts IS 'source of truth for all things table related';

COMMENT ON COLUMN accounts.id IS 'unique account identifier; chosen by user';

CREATE TABLE IF NOT EXISTS events (
  id BIGSERIAL PRIMARY KEY
  , type TEXT NOT NULL
  , data JSONB NOT NULL
);

COMMENT ON TABLE events IS 'log of all events that occur within the system';

COMMENT ON COLUMN events.id IS 'auto-increment unique id';

COMMENT ON COLUMN events.type IS 'what type of event this is';

COMMENT ON COLUMN events.data IS 'information related to the event';

CREATE TABLE IF NOT EXISTS ledger_entries (
  id BIGSERIAL PRIMARY KEY
  , event_id BIGINT NOT NULL REFERENCES events (id)
  , account_id BIGINT NOT NULL REFERENCES accounts (id)
  , amount_micro BIGINT NOT NULL
);

COMMENT ON TABLE ledger_entries IS 'list of changes in balance for each account';

COMMENT ON COLUMN ledger_entries.id IS 'auto-increment unique id';

COMMENT ON COLUMN ledger_entries.account_id IS 'account whose balance is changing';

COMMENT ON COLUMN ledger_entries.amount_micro IS 'micro-amount that balance is changing by - 1 unit = 1,000,000 micro units';

