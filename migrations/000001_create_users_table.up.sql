ALTER DATABASE picpay SET timezone = 'America/Sao_Paulo';
CREATE EXTENSION citext;
CREATE TYPE document_type_enum AS ENUM ('cpf', 'cnpj');
CREATE TYPE balance AS (
    number NUMERIC,
    currency_code TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    name text NOT NULL,
    balance balance NOT NULL,
    document_number text NOT NULL,
    document_type document_type_enum NOT NULL
);