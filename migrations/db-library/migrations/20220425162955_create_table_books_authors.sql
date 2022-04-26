-- +goose Up
-- +goose StatementBegin
CREATE TABLE books_authors (
    book_id     INTEGER NOT NULL,
    author_id   INTEGER NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books_authors;
-- +goose StatementEnd
