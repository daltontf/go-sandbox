package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"main/yascon"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("yascon.db"))
	if err != nil {
	  panic(err)
	}
	//defer db.Close();

	yascon.CreateTablesIfNotExists(db)

  	r := chi.NewRouter()

  	r.Get("/venues", yascon.GetVenues(db))
	r.Post("/venues", yascon.CreateVenue(db))
	r.Put("/venues/{id}", yascon.UpdateVenue(db))
	r.Delete("/venues/{id}", yascon.DeleteVenue(db))

	r.Get("/speakers", yascon.GetSpeakers(db))
	r.Post("/speakers", yascon.CreateSpeaker(db))
	r.Put("/speakers/{id}", yascon.UpdateSpeaker(db))
	r.Delete("/speakers/{id}", yascon.DeleteSpeaker(db))

	r.Get("/session-times", yascon.GetSessionTimes(db))
	r.Post("/session-times", yascon.CreateSessionTime(db))
	r.Put("/session-times/{id}", yascon.UpdateSessionTime(db))
	r.Delete("/session-times/{id}", yascon.DeleteSessionTime(db))

	r.Get("/sessions", yascon.GetSessions(db))
	r.Post("/sessions", yascon.CreateSession(db))
	r.Put("/sessions/{id}", yascon.UpdateSession(db))
	r.Delete("/sessions/{id}", yascon.DeleteSession(db))

	r.Get("/presentations", yascon.GetPresentations(db))
	r.Post("/presentations", yascon.CreatePresentation(db))
	r.Put("/presentations/{id}", yascon.UpdatePresentation(db))
	r.Delete("/presentations/{id}", yascon.DeletePresentation(db))

	r.Get("/attendee-sessions", yascon.GetAttendeeSessions(db))
	r.Post("/attendee-sessions", yascon.CreateAttendeeSession(db))
	r.Put("/attendee-sessions/{id}", yascon.UpdateAttendeeSession(db))
	r.Delete("/attendee-sessions/{id}", yascon.DeleteAttendeeSession(db))

	r.Get("/attendees", yascon.GetAttendees(db))
	r.Post("/attendees", yascon.CreateAttendee(db))
	r.Put("/attendees/{id}", yascon.UpdateAttendee(db))
	r.Delete("/attendees/{id}", yascon.DeleteAttendee(db))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}