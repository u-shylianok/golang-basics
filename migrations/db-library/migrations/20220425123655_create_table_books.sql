-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    edition     INTEGER NOT NULL,
    year        INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
