package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"main/yascon"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

func main() {
	db, err := gorm.Open(sqlite.Open(os.Args[1] + ".db"))
	if err != nil {
	  panic(err)
	}
	//defer db.Close();

	yascon.CreateTablesIfNotExists(db)

  	r := chi.NewRouter()

  	r.Get("/venues", yascon.GetVenues(db))
	r.Get("/venues/{id}", yascon.GetVenue(db))
	r.Post("/venues", yascon.CreateVenue(db))
	r.Put("/venues/{id}", yascon.UpdateVenue(db))
	r.Delete("/venues/{id}", yascon.DeleteVenue(db))

	r.Get("/speakers", yascon.GetSpeakers(db))
	r.Get("/speakers/{id}", yascon.GetSpeaker(db))
	r.Post("/speakers", yascon.CreateSpeaker(db))
	r.Put("/speakers/{id}", yascon.UpdateSpeaker(db))
	r.Delete("/speakers/{id}", yascon.DeleteSpeaker(db))

	r.Get("/sessions", yascon.GetSessions(db))
	r.Get("/sessions/{id}", yascon.GetSession(db))
	r.Post("/sessions", yascon.CreateSession(db))
	r.Put("/sessions/{id}", yascon.UpdateSession(db))
	r.Delete("/sessions/{id}", yascon.DeleteSession(db))

	r.Get("/presentations", yascon.GetPresentations(db))
	r.Get("/presentations/{id}", yascon.GetPresentation(db))
	r.Post("/presentations", yascon.CreatePresentation(db))
	r.Put("/presentations/{id}", yascon.UpdatePresentation(db))
	r.Delete("/presentations/{id}", yascon.DeletePresentation(db))

	r.Get("/attendees", yascon.GetAttendees(db))
	r.Get("/attendees/{id}", yascon.GetAttendee(db))
	r.Post("/attendees", yascon.CreateAttendee(db))
	r.Put("/attendees/{id}", yascon.UpdateAttendee(db))
	r.Delete("/attendees/{id}", yascon.DeleteAttendee(db))

	r.Get("/attendees/{attendees_id}/sessions", yascon.SessionsForAttendee(db))
	r.Get("/sessions/{sessions_id}/attendees", yascon.AttendeesForSession(db))

	r.Put("/attendees/{attendees_id}/sessions/{sessions_id}", yascon.CreateAttendeeSession(db))
	r.Put("/sessions/{sessions_id}/attendees/{attendees_id}", yascon.CreateAttendeeSession(db))
	
	r.Delete("/attendees/{attendees_id}/sessions/{sessions_id}", yascon.DeleteAttendeeSession(db))
	r.Delete("/sessions/{sessions_id}/attendees/{attendees_id}", yascon.DeleteAttendeeSession(db))

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}