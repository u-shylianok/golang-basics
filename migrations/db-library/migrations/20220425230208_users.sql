-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    full_name   VARCHAR(200) NOT NULL,
    passport    TEXT NOT NULL,
    phone       TEXT NOT NULL,
    rating      INTEGER DEFAULT 0 CHECK (rating >= 0 AND rating <= 100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
