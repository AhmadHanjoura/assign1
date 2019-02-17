/*
CMPT 315 (Winter 2018)
Assignment 1
main
Author: Ahmad Hanjoura
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	*Database
}

func (h *Handler) handleGetAllPresenters(w http.ResponseWriter, r *http.Request) {
	presenters, err := h.getAllPresenters()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(presenters)
}

func (h *Handler) handleGetPresenter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	presenter, err := h.getPresenter(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(presenter)
}

func (h *Handler) handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.getQuestions()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(questions)
}

func (h *Handler) handleGetAQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	question, err := h.getAQuestion(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(question)
}

func (h *Handler) handleGetAllTitles(w http.ResponseWriter, r *http.Request) {
	titles, err := h.getAllTitles()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(titles)
}

func (h *Handler) handleInsertAnswer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var answer Answer
	err := decoder.Decode(&answer)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

func main() {
	// connect to database
	connect := "dbname=assign1 user=postgres host=/tmp sslmode=disable"
	db, err := OpenDatabase(connect)
	if err != nil {
		log.Fatal("OpenDatabase %v", err)
	}
	defer db.Close()

	handlers := Handler{
		db,
	}

	router := mux.NewRouter()
	router.Path("/api/v1/presenters").Methods("GET").HandlerFunc(handlers.handleGetAllPresenters)
	router.Path("/api/v1/presenters/{id:[0-9]+}").Methods("GET").HandlerFunc(handlers.handleGetPresenter)
	router.Path("/api/v1/questions").Methods("GET").HandlerFunc(handlers.handleGetQuestions)
	router.Path("/api/v1/questions/{id:[0-9]+}").Methods("GET").HandlerFunc(handlers.handleGetAQuestion)
	router.Path("/api/v1/titles").Methods("GET").HandlerFunc(handlers.handleGetAllTitles)
	router.Path("/api/v1/answers").Methods("POST").HandlerFunc(handlers.handleInsertAnswer)

	log.Fatal(http.ListenAndServe(":8080", router))

}
