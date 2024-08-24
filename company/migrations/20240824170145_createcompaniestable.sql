-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE companies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(15) NOT NULL UNIQUE,
    description VARCHAR(3000),
    amount_of_employees INTEGER NOT NULL,
    registered BOOLEAN NOT NULL,
    type INTEGER NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE companies;
-- +goose StatementEnd
