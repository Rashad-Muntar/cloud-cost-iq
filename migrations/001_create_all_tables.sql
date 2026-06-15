-- +goose Up

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aws_account_id VARCHAR(12) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    environment VARCHAR(20) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS cost_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    service TEXT NOT NULL,
    region TEXT NOT NULL,
    cost_amount NUMERIC(12,4) NOT NULL,
    usage_amount NUMERIC(12,4) NOT NULL,
    currency TEXT DEFAULT 'USD',
    usage_date TIMESTAMP NOT NULL,
    idempotency_key TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS daily_cost_summaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    service TEXT NOT NULL,
    region TEXT NOT NULL,
    total_cost NUMERIC(12,4) NOT NULL,
    total_usage NUMERIC(12,4) NOT NULL,
    currency TEXT DEFAULT 'USD',
    summary_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(account_id, service, region, summary_date)
);

CREATE INDEX idx_cost_events_account_id ON cost_events(account_id);
CREATE INDEX idx_cost_events_usage_date ON cost_events(usage_date);
CREATE INDEX idx_cost_events_account_date ON cost_events(account_id, usage_date DESC);
CREATE INDEX idx_cost_events_service ON cost_events(service);
CREATE INDEX idx_daily_summaries_account_date ON daily_cost_summaries(account_id, summary_date DESC);

-- +goose Down
DROP TABLE IF EXISTS daily_cost_summaries;
DROP TABLE IF EXISTS cost_events;
DROP TABLE IF EXISTS accounts;