-- +goose Up
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aws_account_id VARCHAR(12) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    environment VARCHAR(20) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes for performance
CREATE INDEX idx_accounts_aws_id ON accounts(aws_account_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_accounts_environment ON accounts(environment) WHERE deleted_at IS NULL;

-- +goose Down
DROP TABLE IF EXISTS accounts;