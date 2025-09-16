package yascon

import (
	"context"
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func getAll[T any](db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		venues, err := gorm.G[T](db).Find(ctx)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(venues)
	}	
}

func getById[T any](db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		ctx := context.Background()

		entity, err := gorm.G[T](db).Where("id = ?", id).First(ctx)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(entity)
		w.WriteHeader(http.StatusOK)
	}
}

func create[T HasID](db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res T
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		ctx := context.Background()

		err := gorm.G[T](db).Create(ctx, &res)
		if err != nil {
			http.Error(w, err.Error(), 409)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s/%d", r.URL.Path, res.GetID()))
		w.WriteHeader(http.StatusCreated)
	}
}

func deleteById[T any](db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		ctx := context.Background()

		_, err := gorm.G[T](db).Where("id = ?", id).Delete(ctx)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func updateById[T any](db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var res T
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		ctx := context.Background()

		_, err := gorm.G[T](db).Where("id = ?", id).Updates(ctx, res)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetVenues(db *gorm.DB) http.HandlerFunc {
	return getAll[Venue](db)
}

func GetVenue(db *gorm.DB) http.HandlerFunc {
	return getById[Venue](db)
}

func CreateVenue(db *gorm.DB) http.HandlerFunc {
	return create[Venue](db)
}

func DeleteVenue(db *gorm.DB) http.HandlerFunc {
	return deleteById[Venue](db)
}

func UpdateVenue(db *gorm.DB) http.HandlerFunc {
	return updateById[Venue](db)
}


func GetSpeakers(db *gorm.DB) http.HandlerFunc {
	return getAll[Speaker](db)
}

func GetSpeaker(db *gorm.DB) http.HandlerFunc {
	return getById[Speaker](db)
}

func CreateSpeaker(db *gorm.DB) http.HandlerFunc {
	return create[Speaker](db)
}

func DeleteSpeaker(db *gorm.DB) http.HandlerFunc {
	return deleteById[Speaker](db)
}

func UpdateSpeaker(db *gorm.DB) http.HandlerFunc {
	return updateById[Speaker](db)
}

func GetSessions(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		presentationFields := r.URL.Query()["presentation"]
		venueFields := r.URL.Query()["venue"]
		if len(presentationFields) == 0 || len(venueFields) == 0 {
			getAll[Presentation](db)(w, r)
		} else {
			GetSessionsWithPresentationAndVenue(db)(w, r)
		}
	}
}

func GetSession(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		presentationFields := r.URL.Query()["presentation"]
		venueFields := r.URL.Query()["venue"]
		if len(presentationFields) == 0 || len(venueFields) == 0 {
			getAll[Presentation](db)(w, r)
		} else {
			GetSessionWithPresentationAndVenue(db)(w, r)
		}
	}
}

func CreateSession(db *gorm.DB) http.HandlerFunc {
	return create[Session](db)
}

func DeleteSession(db *gorm.DB) http.HandlerFunc {
	return deleteById[Session](db)
}

func UpdateSession(db *gorm.DB) http.HandlerFunc {
	return updateById[Session](db)
}

func GetSessionsWithPresentationAndVenue(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sessionsWithPresentationAndVenue []SessionWithPresentationAndVenue
		// TODO handle more than assumeing presentation and venue 
		db.Model(&Session{}).
		Select("sessions.*, presentations.id as presentation_id, presentations.name as presentation_name, venues.id as venue_id, venues.name as venue_name").
		Joins("JOIN presentations ON sessions.presentations_id = presentations.id").
		Joins("JOIN venues ON sessions.venues_id = venues.id").
    		Scan(&sessionsWithPresentationAndVenue)

		w.Header().Set("Content-Type", "application/json")

		if sessionsWithPresentationAndVenue != nil {
			json.NewEncoder(w).Encode(sessionsWithPresentationAndVenue)
		}
	}
}

func GetSessionWithPresentationAndVenue(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sessionWithPresentationAndVenue SessionWithPresentationAndVenue
		// TODO handle more than assumeing presentation and venue 
		error := db.Model(&Session{}).
    		Select("sessions.*, presentations.id as presentation_id, presentations.name as presentation_name, venues.id as venue_id, venues.name as venue_name").
    		Joins("JOIN presentations ON sessions.presentations_id = presentations.id").
			Joins("JOIN venues ON sessions.venues_id = venues.id").
    		First(&sessionWithPresentationAndVenue)

		w.Header().Set("Content-Type", "application/json")

		if error != nil {
			json.NewEncoder(w).Encode(sessionWithPresentationAndVenue)
		}
	}
}


func GetPresentationsWithSpeaker(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var presentationsWithSpeaker []PresentationWithSpeaker
		// TODO handle more than assumeing name 
		db.Model(&Presentation{}).
			Select("presentations.*, speakers.id as speaker_id, speakers.name as speaker_name").
			Joins("JOIN speakers ON presentations.speakers_id = speakers.id").
    		Scan(&presentationsWithSpeaker)

		w.Header().Set("Content-Type", "application/json")

		if presentationsWithSpeaker != nil {
			json.NewEncoder(w).Encode(presentationsWithSpeaker)
		}
		
	}
}

func GetPresentationWithSpeaker(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var presentationWithSpeaker PresentationWithSpeaker
		// TODO handle more than assumeing name 
		error := db.Model(&Presentation{}).
    		Select("presentations.*, speakers.id as speaker_id, speakers.name as speaker_name").
    		Joins("JOIN speakers ON presentations.speakers_id = speakers.id").
    		First(&presentationWithSpeaker)

		w.Header().Set("Content-Type", "application/json")

		if error != nil {
			json.NewEncoder(w).Encode(presentationWithSpeaker)
		}
	}
}

func GetPresentations(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		speakerFields := r.URL.Query()["speaker"]
		if len(speakerFields) == 0 {
			getAll[Presentation](db)(w, r)
		} else {
			GetPresentationsWithSpeaker(db)(w, r)
		}
	}
}

func GetPresentation(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		speakerFields := r.URL.Query()["speaker"]
		if len(speakerFields) == 0 {
			getById[Presentation](db)(w, r)
		} else {
			GetPresentationWithSpeaker(db)(w, r)
		}
	} 
}

func CreatePresentation(db *gorm.DB) http.HandlerFunc {
	return create[Presentation](db)
}

func DeletePresentation(db *gorm.DB) http.HandlerFunc {
	return deleteById[Presentation](db)
}

func UpdatePresentation(db *gorm.DB) http.HandlerFunc {
	return updateById[Presentation](db)
}

func CreateAttendeeSession(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		attendeeId, _ :=  strconv.Atoi(chi.URLParam(r, "attendees_id"))
		sessionId, _ := strconv.Atoi(chi.URLParam(r, "sessions_id"))
		res := AttendeeSession {
			Attendees_ID: attendeeId,
			Sessions_ID: sessionId,
		}

		ctx := context.Background()

		err := gorm.G[AttendeeSession](db).Create(ctx, &res)
		if err != nil {
			http.Error(w, err.Error(), 409)
			return
		}
		w.Header().Set("Location", r.URL.Path)
		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteAttendeeSession(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		attendeesId := chi.URLParam(r, "attendees_id")
		sessionsId := chi.URLParam(r, "sessions_id")

		ctx := context.Background()

		_, err := gorm.G[AttendeeSession](db).Where("attendees_id = ? and sessions_id = ?", attendeesId, sessionsId).Delete(ctx)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func SessionsForAttendee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		attendeesId :=  chi.URLParam(r, "attendees_id")

		attendeeSessions, err := gorm.G[AttendeeSession](db).Where("attendees_id = ?", attendeesId).Find(ctx)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(attendeeSessions)
	}	
}

func AttendeesForSession(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		sessionsId := chi.URLParam(r, "sessions_id")

		attendeeSessions, err := gorm.G[AttendeeSession](db).Where("sessions_id = ?", sessionsId).Find(ctx)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(attendeeSessions)
	}	
}


func GetAttendees(db *gorm.DB) http.HandlerFunc {
	return getAll[Attendee](db)
}

func GetAttendee(db *gorm.DB) http.HandlerFunc {
	return getById[Attendee](db)
}

func CreateAttendee(db *gorm.DB) http.HandlerFunc {
	return create[Attendee](db)
}

func DeleteAttendee(db *gorm.DB) http.HandlerFunc {
	return deleteById[Attendee](db)
}

func UpdateAttendee(db *gorm.DB) http.HandlerFunc {
	return updateById[Attendee](db)
}