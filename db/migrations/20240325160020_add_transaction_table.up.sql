CREATE TABLE IF NOT EXISTS transactions (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id) NOT NULL,
  bank_id UUID REFERENCES bank_accounts(id) NOT NULL,
  currency VARCHAR(3) NOT NULL,
  balance int NOT NULL,
  image_url VARCHAR(255) DEFAULT '',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_transactions ON transactions (currency, created_at)