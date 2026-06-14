CREATE TABLE recommendations (

    id UUID PRIMARY KEY,

    account_id UUID NOT NULL,

    type TEXT NOT NULL,

    title TEXT NOT NULL,

    description TEXT NOT NULL,

    estimated_monthly_savings NUMERIC(14,2)
    NOT NULL,

    severity TEXT NOT NULL,

    status TEXT NOT NULL,

    created_at TIMESTAMP
    NOT NULL
    DEFAULT NOW(),

    CONSTRAINT fk_rec_account
    FOREIGN KEY(account_id)
    REFERENCES accounts(id)
);

CREATE INDEX idx_rec_account
ON recommendations(account_id);