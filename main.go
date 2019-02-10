// CMPT 315 (Winter 2018)
//
// Assignment 1
// main
// Author: Ahmad Hanjoura

package main

import (
	"log"
	"net/http"
)

func (h *Handler) handleGetAllPresenters(w http.ResponseWriter, r *http.Request) {

}

func main() {
	connect := "dbname=assign1 user=postgres host=localhost port=5432 sslmode=disable"
	db, err := OpenDatabase(connect)
	if err != nil {
		log.Fatal("OpenDatabase %v", err)
	}
	defer db.Close()

}
