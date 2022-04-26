-- +goose Up
-- +goose StatementBegin
CREATE TABLE electronic_orders (
    id      SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    date    TIMESTAMP NOT NULL DEFAULT NOW(),
    price   MONEY NOT NULL,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE electronic_orders;
-- +goose StatementEnd
