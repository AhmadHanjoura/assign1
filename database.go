/*
 CMPT 315 (Winter 2018)
 Assignment 1
 SQL queries for postgreSQL database
 Author: Ahmad Hanjoura
*/

package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//-------------STRUCTS----------------

type Database struct {
	*sqlx.DB
}

// struct user represents the data stored in the
// table for a single user
type User struct {
	UserID    int    `db:"user_id" json:"userId"` // json:"userId, omitempty"  if returning it with a missing field
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
}

// struct presentation represents the data stored in the
// table for a single presentation and presenter
type Presentation struct {
	PresentationID int       `db:"pid" json:"pid,omitempty"`
	PresenterID    int       `db:"presenter_id" json:"presenterId,omitempty"`
	FirstName      string    `db:"first_name" json:"firstName,omitempty"`
	LastName       string    `db:"last_name" json:"lastName,omitempty"`
	Title          string    `db:"title" json:"title"`
	Date           time.Time `db:"p_date" json:"date,omitempty"`
}

// struct question represents the data stored in the
// table for a single question
type Question struct {
	QuestionID int    `db:"qid"`
	Type       string `db:"q_type"` // (MC1) OR (OPEN2)
	Question   string `db:"question"`
}

// struct answer represents the data stored in the
// table for a single answer
type Answer struct {
	AnswerID    int    `db:"aid" json:"aid,omitempty"`
	QuestionID  int    `db:"qid" json:"questionId"`
	PresenterID int    `db:"presenter_id" json:"presenterId"`
	ReviewerID  int    `db:"reviewer_id" json:"reviewerId"`
	Answer      string `db:"answer" json:"answer"`
}

// Attempts to open database
// Sourced from lab04
func OpenDatabase(connect string) (*Database, error) {
	db := Database{}
	var err error

	db.DB, err = sqlx.Connect("postgres", connect)
	if err != nil {
		return nil, fmt.Errorf("Connect (%v): %v", connect, err)
	}

	return &db, nil
}

// Prints the presenter information
// Adapted from Lab03 solution
func printPresenters(presenters []Presentation) {
	fmt.Printf("ID   Last name, First name, Title\n")
	fmt.Printf("--------------------------------\n")
	for _, presenter := range presenters {
		fmt.Printf("%-4d %s, %s, %s\n", presenter.PresenterID, presenter.LastName, presenter.FirstName, presenter.Title)
	}
}

//------------QUERIES-------------

func (db *Database) getAllPresenters() ([]Presentation, error) {
	query := `SELECT presenter_id, first_name, last_name, title 
				FROM presenters, users, presentations
				WHERE presenters.user_id = users.user_id 
				AND presenters.pid = presentations.pid`
	presenters := []Presentation{}

	err := db.Select(&presenters, query)
	if err != nil {
		return nil, err
	}

	return presenters, err
}

func (db *Database) getPresenter(presenterID string) (Presentation, error) {
	query := `SELECT title, first_name, last_name, p_date
	FROM presenters, users, presentations
	WHERE presenter_id = $1 AND presenters.user_id = users.user_id AND presenters.pid = presentations.pid`

	presenter := Presentation{}

	err := db.Get(&presenter, query, presenterID)
	if err != nil {
		return presenter, err
	}

	return presenter, err
}

func (db *Database) getQuestions() ([]Question, error) {
	query := `SELECT * FROM questions`

	questions := []Question{}

	err := db.Select(&questions, query)
	if err != nil {
		return nil, err
	}

	return questions, err
}

func (db *Database) getAQuestion(qid string) (Question, error) {
	query := `SELECT * FROM questions WHERE qid = $1`

	question := Question{}

	err := db.Get(&question, query, qid)
	if err != nil {
		return question, err
	}

	return question, err
}

func (db *Database) getAllTitles() ([]Presentation, error) {
	query := `SELECT title FROM presentations`

	titles := []Presentation{}

	err := db.Select(&titles, query)
	if err != nil {
		return nil, err
	}

	return titles, err
}

func (db *Database) insertAnswer(qid, presenterID, reviewerID, answer string) error {
	query := `INSERT INTO answers(qid, presenter_id, reviewer_id, answer)
				VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, qid, presenterID, reviewerID, answer)
	if err != nil {
		return err
	}

	return err
}

/*func (db *Database) giveAnswer() error {
	query := `INSERT INTO answers
				VALUES()`
}*/
