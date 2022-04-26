-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id              SERIAL PRIMARY KEY,
    book_copy_id    INTEGER NOT NULL,
    user_id         INTEGER NOT NULL,
    librarian_id    INTEGER NOT NULL,
    start_date      TIMESTAMP NOT NULL DEFAULT NOW(),
    end_date        TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '1 MONTH' * 3,
    is_returned     BOOLEAN NOT NULL,
    start_condition INTEGER NOT NULL CHECK (start_condition >= 0 AND start_condition <= 100),
    end_condition   INTEGER NOT NULL CHECK (start_condition >= 0 AND start_condition <= 100),
    price           MONEY NOT NULL,
    FOREIGN KEY (book_copy_id) REFERENCES book_copies(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (librarian_id) REFERENCES librarians(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
