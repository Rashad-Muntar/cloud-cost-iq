-- +goose Up
CREATE TABLE IF NOT EXISTS cost_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id TEXT NOT NULL,
    service TEXT NOT NULL,
    region TEXT NOT NULL,
    cost_amount NUMERIC(12,4) NOT NULL,
    usage_amount NUMERIC(12,4) NOT NULL,
    currency TEXT DEFAULT 'USD',
    usage_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);