// CMPT 315 (Winter 2018)
//
// Assignment 1
// SQL queries for postgreSQL database
// Author: Ahmad Hanjoura

package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

//-------------STRUCTS----------------

type Database struct {
	*sqlx.DB
}

// struct user represents the data stored in the
// table for a single user
type User struct {
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// struct presentation represents the data stored in the
// table for a single presentation and presenter
type Presentation struct {
	PresentationID int       `db:"pid"`
	PresenterID    int       `db:"presenter_id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Title          string    `db:"title"`
	Date           time.Time `db:"p_date"`
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
	AnswerID    int    `db:"aid"`
	QuestionID  int    `db:"qid"`
	PresenterID int    `db:"presenter_id"`
	ReviewerID  int    `db:"reviewer_id"`
	Answer      string `db:"answer"`
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

func (db *Database) getAllPresenters(db *sqlx.DB) error {
	query := `SELECT presenter_id, first_name, last_name, title 
				FROM presenters, users, presentations
				WHERE presenters.user_id = users.user_id 
				AND presenters.pid = presentations.pid`
	presenters := []Presentation{}

	err := db.Select(&presenters, query)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) getPresenter(db *sqlx.DB) error {
	query := `SELECT title, first_name, last_name, p_date
	FROM presenters, users, presentations
	WHERE presenters.user_id = users.user_id AND presenters.pid = presentations.pid 
	AND first_name = $1 AND last_name = $2`

	presenter := Presentation{}

	err := db.Get(&presenter, query, presenter.FirstName, presenter.LastName)
	if err != nil {
		return err
	}

	return nil
}
