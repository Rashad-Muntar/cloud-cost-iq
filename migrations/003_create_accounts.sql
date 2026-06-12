-- +goose Up
CREATE TABLE IF NOT EXISTS accounts (
    internal_id UUID PRIMARY KEY,
    aws_account_id VARCHAR(12) UNIQUE NOT NULL,
    organizational_path TEXT NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL,
    environment VARCHAR(20) NOT NULL,
    owner_team VARCHAR(255) NOT NULL DEFAULT '',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    tags JSONB DEFAULT '{}',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    business_unit VARCHAR(255) NOT NULL DEFAULT '',
    cost_center VARCHAR(50) NOT NULL DEFAULT '',
    is_management_account BOOLEAN DEFAULT FALSE,
    region_constraint TEXT[] DEFAULT '{}'
);

-- Indexes for performance
CREATE INDEX idx_accounts_aws_id ON accounts(aws_account_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_accounts_environment ON accounts(environment) WHERE deleted_at IS NULL;
CREATE INDEX idx_accounts_status ON accounts(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_accounts_owner_team ON accounts(owner_team) WHERE deleted_at IS NULL;
CREATE INDEX idx_accounts_business_unit ON accounts(business_unit) WHERE deleted_at IS NULL;

-- Comments for documentation
COMMENT ON TABLE accounts IS 'Internal representation of AWS accounts';
COMMENT ON COLUMN accounts.internal_id IS 'Our internal UUID, not exposed to users';
COMMENT ON COLUMN accounts.aws_account_id IS '12-digit AWS account ID, immutable';
COMMENT ON COLUMN accounts.status IS 'active|disabled|archived|pending';