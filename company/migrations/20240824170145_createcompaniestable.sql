-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies (
    id UUID PRIMARY KEY,
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
