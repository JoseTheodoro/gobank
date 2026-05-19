CREATE TABLE onboarding_processes (
    id BIGSERIAL PRIMARY KEY,
    onboarding_id uuid NOT NULL UNIQUE,
    customer_id uuid NULL,
    account_id uuid NULL,
    email varchar(255) NOT NULL UNIQUE,
    document varchar(16) NOT NULL UNIQUE ,
    status varchar(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);