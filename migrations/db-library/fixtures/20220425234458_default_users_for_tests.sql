-- +goose Up
-- +goose StatementBegin
INSERT INTO users(full_name, passport, phone, rating) VALUES
    ('BodyGueard Ежов', 'AA1111111', '+1 (123) 1231231', 100),
    ('Пёрсон Фулл Нейм', 'AA1111112', '+1 (123) 3213213', 15);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE full_name = 'BodyGueard Ежов' OR full_name = 'Пёрсон Фулл Нейм';
-- +goose StatementEnd
