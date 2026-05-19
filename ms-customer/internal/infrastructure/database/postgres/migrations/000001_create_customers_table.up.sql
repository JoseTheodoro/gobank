CREATE TABLE customers (
    id BIGSERIAL PRIMARY KEY,
    customer_id uuid NOT NULL UNIQUE,
    name varchar(100) NOT NULL,
    email varchar (100) NOT NULL,
    type varchar (50) NOT NULL,
    document varchar(16) NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NULL
)