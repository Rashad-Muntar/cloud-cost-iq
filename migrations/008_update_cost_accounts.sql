ALTER TABLE cost_events
ADD COLUMN idempotency_key TEXT UNIQUE;