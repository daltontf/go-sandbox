package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
  	"github.com/glebarez/sqlite"
  	"gorm.io/gorm"
)

func main() {
 	db, err := gorm.Open(sqlite.Open("learning.db"))
  	if err != nil {
    	panic(err)
  	}
	db.Exec(`CREATE TABLE IF NOT EXISTS resources (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		url TEXT NOT NULL,
		status TEXT NOT NULL);`)

  	r := chi.NewRouter()

  	r.Get("/resources", getResources(db))
	r.Post("/resource", createResource(db))
	r.Put("/resource/{id}", updateResource(db))
	r.Delete("/resource/{id}", deleteResource(db))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}