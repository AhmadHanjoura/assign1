-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL
);

CREATE TABLE presentations (
    pid SERIAL PRIMARY KEY,
    title text NOT NULL,
    p_date timestamp NOT NULL
);

CREATE TABLE questions (
    qid SERIAL PRIMARY KEY,
    q_type text NOT NULL,
    question text NOT NULL
);

CREATE TABLE presenters (
    presenter_id SERIAL PRIMARY KEY,
    user_id integer REFERENCES users(user_id),
    pid integer REFERENCES presentations(pid)
);

CREATE TABLE answers (
    aid SERIAL PRIMARY KEY,
    qid integer REFERENCES questions(qid),
    presenter_id integer REFERENCES presenters(presenter_id),
    reviewer_id integer REFERENCES users(user_id),
    answer text
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
DROP TABLE presentations;
DROP TABLE questions;
DROP TABLE answers;
DROP TABLE presenters;