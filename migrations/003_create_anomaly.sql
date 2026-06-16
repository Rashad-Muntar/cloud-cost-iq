-- +goose Up
CREATE TABLE anomalies (
    id UUID PRIMARY KEY,
    account_id UUID NOT NULL,
    service TEXT NOT NULL,
    expected_cost NUMERIC(14,2) NOT NULL,
    actual_cost NUMERIC(14,2) NOT NULL,
    deviation NUMERIC(14,2) NOT NULL,
    severity TEXT NOT NULL,
    detected_at TIMESTAMP NOT NULL
    DEFAULT NOW(),
    CONSTRAINT fk_anomaly_account
    FOREIGN KEY(account_id)
    REFERENCES accounts(id)
);

CREATE INDEX idx_anomaly_account
ON anomalies(account_id);