// CMPT 315 (Winter 2018)
//
// Assignment 1
// SQL queries for postgreSQL database
// Author: Ahmad Hanjoura

package main

import (
	"time"
)

// struct user represents the data stored in the
// table for a single user
type user struct {
	UserID    int    `db:"uid"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// struct presentation represents the data stored in the
// table for a single presentation
type presentation struct {
	PresentationID int       `db:"pid"`
	Title          string    `db:"title"`
	Date           time.Time `db:"date"`
}

// struct question represents the data stored in the
// table for a single question
type question struct {
	QuestionID int    `db:"qid"`
	Type              // FINISH THIS. MAYBE INCLUDE QUESTION NUMBERS LIKE (MC1) OR (OPEN2)
	Question   string `db:"question"`
}

// struct answer represents the data stored in the
// table for a single answer
type answer struct {
	AnswerID   int `db:"aid"`
	QuestionID int `db:"qid"`
	Type
	Answer string `db:"answer"`
}

// struct presenter represents the data stored in the
// table for a single presenter
type presenter struct {
	UserID         int `db:"uid"`
	PresentationID int `db:"pid"`
}
