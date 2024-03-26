CREATE TABLE IF NOT EXISTS bank_accounts (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id) NOT NULL,
  account_name varchar(30) NOT NULL,
  account_number varchar(30) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);