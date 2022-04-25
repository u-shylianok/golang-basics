-- +goose Up
-- +goose StatementBegin
CREATE TABLE librarians (
    id          SERIAL PRIMARY KEY,
    full_name   VARCHAR(200) NOT NULL,
    position    INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE librarians;
-- +goose StatementEnd
