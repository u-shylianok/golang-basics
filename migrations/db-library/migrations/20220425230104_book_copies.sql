-- +goose Up
-- +goose StatementBegin
CREATE TABLE book_copies (
    id              SERIAL PRIMARY KEY,
    book_id         INTEGER NOT NULL,
    copy_type       VARCHAR(50) NOT NULL,
    start_condition INTEGER DEFAULT 100 CHECK (start_condition >= 0 AND start_condition <= 100),
    FOREIGN KEY (book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book_copies;
-- +goose StatementEnd
