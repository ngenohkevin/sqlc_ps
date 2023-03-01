CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname  TEXT NOT NULL
);

CREATE TABLE todos (
    id BIGSERIAL PRIMARY KEY,
    users_id BIGSERIAL NOT NULL REFERENCES users(id),
    task TEXT NOT NULL,
    done BOOLEAN NOT NULL
);