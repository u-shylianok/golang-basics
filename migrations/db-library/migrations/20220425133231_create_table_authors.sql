-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors (
    id              SERIAL PRIMARY KEY,
    full_name       VARCHAR(200) NOT NULL,
    year_of_birth   INTEGER NOT NULL,
    description     TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
-- +goose StatementEnd
