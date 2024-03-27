CREATE TABLE IF NOT EXISTS bank_accounts (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id) NOT NULL,
  currency VARCHAR(3) NOT NULL,
  total_balance int NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_currency ON bank_accounts (currency);

CREATE INDEX IF NOT EXISTS idx_balance ON bank_accounts (total_balance);