-- +goose Up
-- +goose StatementBegin
CREATE TYPE book_availability AS ENUM ('PHYSICAL_ONLY', 'VIRTUAL_ONLY', 'PHYSICAL_AND_VIRTUAL');

ALTER TABLE books
    ADD availability book_availability  NOT NULL,
    ADD physical_copies_available       INTEGER NOT NULL,
    ADD rent_price                      MONEY NOT NULL,
    ADD e_copy_price                    MONEY NOT NULL;

CREATE TYPE order_type AS ENUM ('RENT', 'E_COPY_BOOKING');

ALTER TABLE orders
    ADD COLUMN book_id  INTEGER NOT NULL,
    ADD COLUMN type     order_type NOT NULL,
    ADD CONSTRAINT orders_book_id_fkey FOREIGN KEY (book_id) REFERENCES books (id),
    ALTER COLUMN end_date           DROP NOT NULL,
    ALTER COLUMN is_returned        DROP NOT NULL,
    ALTER COLUMN start_condition    DROP NOT NULL,
    ALTER COLUMN end_condition      DROP NOT NULL,
    DROP CONSTRAINT orders_book_copy_id_fkey,
    DROP COLUMN book_copy_id;

DROP TABLE book_copies;
DROP TABLE electronic_orders;

-- +goose StatementEnd

-- +goose Down
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

CREATE TABLE book_copies (
    id              SERIAL PRIMARY KEY,
    book_id         INTEGER NOT NULL,
    copy_type       VARCHAR(50) NOT NULL,
    start_condition INTEGER DEFAULT 100 CHECK (start_condition >= 0 AND start_condition <= 100),
    FOREIGN KEY (book_id) REFERENCES books(id)
);

ALTER TABLE orders
    ADD COLUMN book_copy_id         INTEGER NOT NULL,
    ADD CONSTRAINT orders_book_copy_id_fkey FOREIGN KEY (book_copy_id) REFERENCES book_copies (id),
    ALTER COLUMN end_date           SET NOT NULL,
    ALTER COLUMN is_returned        SET NOT NULL,
    ALTER COLUMN start_condition    SET NOT NULL,
    ALTER COLUMN end_condition      SET NOT NULL,
    DROP CONSTRAINT orders_book_id_fkey,
    DROP COLUMN book_id,
    DROP COLUMN type;

DROP TYPE order_type;

ALTER TABLE books
    DROP COLUMN availability,
    DROP physical_copies_available,
    DROP rent_price,
    DROP e_copy_price;

DROP TYPE book_availability;
-- +goose StatementEnd
