-- +goose Up
CREATE TABLE users (
    uid int is NOT NULL,
    first_name text is NOT NULL,
    last_name text is NOT NULL,
    PRIMARY KEY(uid)
);



-- +goose Down
DROP TABLE users;