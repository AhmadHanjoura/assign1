-- +goose Up
CREATE TABLE users (
    uid SERIAL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL
);

CREATE TABLE presentations (
    pid SERIAL PRIMARY KEY,
    title text NOT NULL,
    date timestamp NOT NULL
);

CREATE TABLE questions (
    qid SERIAL PRIMARY KEY,
    type text NOT NULL,
    question text NOT NULL
);

CREATE TABLE answers (
    aid SERIAL PRIMARY KEY,
    qid integer REFERENCES questions(qid),
    type text REFERENCES questions(type),
    answer text
);

CREATE TABLE presenters (
    uid integer REFERENCES users(uid),
    pid integer REFERENCES presentations(pid)
);


-- +goose Down
DROP TABLE users;
DROP TABLE presentations;
DROP TABLE questions;
DROP TABLE answers;
DROP TABLE presenters;