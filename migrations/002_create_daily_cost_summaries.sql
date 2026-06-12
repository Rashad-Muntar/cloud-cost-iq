-- +goose Up
CREATE TABLE IF NOT EXISTS daily_cost_summaries (
    summary_date DATE NOT NULL,
    service      TEXT NOT NULL,

    total_cost NUMERIC(14,4) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (
        summary_date,
        service
    )
);