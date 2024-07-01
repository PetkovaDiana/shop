-- +goose Up
-- +goose StatementBegin
CREATE TABLE client
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    number BIGINT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL
);

CREATE TABLE category
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL
);

CREATE TABLE product
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL,
    category_id BIGINT NOT NULL,

    CONSTRAINT category_id_fk FOREIGN KEY (category_id)
        REFERENCES category(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product, category, client;
-- +goose StatementEnd
