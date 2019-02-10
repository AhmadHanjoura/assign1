-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO users(first_name, last_name) VALUES('Nicholas', 'Boers');
INSERT INTO users(first_name, last_name) VALUES('Ahmad', 'Hanjoura');
INSERT INTO users(first_name, last_name) VALUES('Jamie', 'Rajewski');
INSERT INTO users(first_name, last_name) VALUES('Lukas', 'Jenks');
INSERT INTO users(first_name, last_name) VALUES('Ahmad', 'Hanjoura');

INSERT INTO presentations(title, p_date) VALUES('Spider-man Fun Facts', '2019-07-08 12:21:13');
INSERT INTO presentations(title, p_date) VALUES('1000 Ways to Die in Dark Souls', '2019-04-16 14:45:12');
INSERT INTO presentations(title, p_date) VALUES('The Legend of Link: Zeldas Turn', '2019-07-04 12:34:56');
INSERT INTO presentations(title, p_date) VALUES('The Batman', '2019-07-08 20:20:20');

INSERT INTO questions(q_type, question) VALUES('mc1', 'Preparedness: the presenter was adequately prepared.');
INSERT INTO questions(q_type, question) VALUES('mc2', 'Organization: the presentation material was arranged logically.');
INSERT INTO questions(q_type, question) VALUES('mc3', 'Correctness: the presented facts were correct (to the best of your knowledge).');
INSERT INTO questions(q_type, question) VALUES('mc4', 'Visualization: the visual material included appropriate content/fonts/graphics.');
INSERT INTO questions(q_type, question) VALUES('mc5', 'General introduction: the presentation clearly introduced the broad area containing the topic.');
INSERT INTO questions(q_type, question) VALUES('mc6', 'Motivation: the presentation clearly motivated the specific topic in the context of the broad area.');
INSERT INTO questions(q_type, question) VALUES('mc7', 'Introduction: the presentation clearly introduced the specific topic.');
INSERT INTO questions(q_type, question) VALUES('mc8', 'Tutorial/demonstration: the tutorial/demonstration improved your understanding of the specific topic.');
INSERT INTO questions(q_type, question) VALUES('mc9', 'Multiple-choice questions: at least three multiple-choice questions assessed your understanding of the presented content.');
INSERT INTO questions(q_type, question) VALUES('mc10', 'Answers: the presenters answers to questions were satisfying.');
INSERT INTO questions(q_type, question) VALUES('open1', 'Provide any comments for the presenter.');
INSERT INTO questions(q_type, question) VALUES('open2', 'Provide any comments for your instructor.');

INSERT INTO presenters(user_id, pid) VALUES(2, 1);
INSERT INTO presenters(user_id, pid) VALUES(3, 2);
INSERT INTO presenters(user_id, pid) VALUES(4, 3);
INSERT INTO presenters(user_id, pid) VALUES(5, 4);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM users;
DELETE FROM presentations;
DELETE FROM questions;
DELETE FROM answers;
DELETE FROM presenters;