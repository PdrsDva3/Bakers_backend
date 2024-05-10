-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id              SERIAL NOT NULL PRIMARY KEY,
    phone           BIGINT UNIQUE,
    hashed_password VARCHAR,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS admin(
    id              SERIAL NOT NULL PRIMARY KEY,
    phone           BIGINT UNIQUE,
    hashed_password VARCHAR
);

CREATE TABLE IF NOT EXISTS bread(
    id          SERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR,
    price       NUMERIC,
    description VARCHAR,
    count       INTEGER,
    photo       bytea
);

CREATE TABLE IF NOT EXISTS orders(
    id      SERIAL NOT NULL PRIMARY KEY,
    id_user INTEGER REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS order_bread_link(
    id_order INTEGER REFERENCES orders(id),
    id_bread INTEGER REFERENCES bread(id)
);

CREATE TABLE IF NOT EXISTS courier(
    id              SERIAL NOT NULL PRIMARY KEY,
    name            VARCHAR,
    phone           BIGINT UNIQUE,
    hashed_password VARCHAR
);

CREATE TABLE IF NOT EXISTS order_courier_link(
    id_order   INTEGER REFERENCES orders(id),
    id_courier INTEGER REFERENCES courier (id)
);

CREATE TABLE IF NOT EXISTS chats(
    id   SERIAL PRIMARY KEY,
    name VARCHAR UNIQUE
);

CREATE TABLE IF NOT EXISTS messages(
    id             SERIAL PRIMARY KEY,
    sender         VARCHAR,
    content        VARCHAR,
    send_timestamp timestamp,
    chat_id        INTEGER REFERENCES chats(id)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
