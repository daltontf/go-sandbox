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

func getAll[T any](db *gorm.DB, ordering *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		var entities []T
		var err error

		if (ordering != nil) {
			entities, err =  gorm.G[T](db).Order(*ordering).Find(ctx)
		} else {
			entities, err =  gorm.G[T](db).Find(ctx)
		}		
		
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(entities)
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
		//w.WriteHeader(http.StatusOK)
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

func stringPtr(s string) *string { return &s }

func GetVenues(db *gorm.DB) http.HandlerFunc {
	return getAll[Venue](db, stringPtr("venues.name asc"))
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
	return getAll[Speaker](db, stringPtr("speakers.name asc"))
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
			getAll[Session](db, stringPtr("sessions.start_time_min asc"))(w, r)
		} else {
			getAll[SessionWithPresentationAndVenue](db, stringPtr("start_time_min asc, presentation_name asc"))(w, r)
		}
	}
}

func GetSession(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		presentationFields := r.URL.Query()["presentation"]
		venueFields := r.URL.Query()["venue"]
		if len(presentationFields) == 0 || len(venueFields) == 0 {
			getById[Session](db)(w, r)
		} else {
			getById[SessionWithPresentationAndVenue](db)(w, r)
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

func GetPresentations(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		speakerFields := r.URL.Query()["speaker"]
		if len(speakerFields) == 0 {
			getAll[Presentation](db, stringPtr("presentations.name asc"))(w, r)
		} else {
			getAll[PresentationWithSpeaker](db, stringPtr("name asc"))(w, r)
		}
	}
}

func GetPresentation(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		speakerFields := r.URL.Query()["speaker"]
		if len(speakerFields) == 0 {
			getById[Presentation](db)(w, r)
		} else {
			getById[PresentationWithSpeaker](db)(w, r)
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
		attendeesId :=  chi.URLParam(r, "attendees_id")
		
		sessionsWithPresentationAndVenue := []SessionWithPresentationAndVenue{}
		
		result := db.Model(&Session{}).
			Select("sessions.*, presentations.id as presentation_id, presentations.name as presentation_name, venues.id as venue_id, venues.name as venue_name").
			Joins("JOIN attendee_sessions ON attendee_sessions.sessions_id = sessions.id").
			Joins("JOIN presentations ON sessions.presentations_id = presentations.id").
			Joins("JOIN venues ON sessions.venues_id = venues.id").
			Order("sessions.start_time_min asc").
			Order("presentations.name asc").
			Where("attendee_sessions.attendees_id = ?", attendeesId).			
    		Scan(&sessionsWithPresentationAndVenue)


		w.Header().Set("Content-Type", "application/json")
		if result.Error == nil {
			json.NewEncoder(w).Encode(sessionsWithPresentationAndVenue)
		} else {
			http.Error(w, result.Error.Error(), 500)
			return	
		}
	}	
}

func AttendeesForSession(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionsId := chi.URLParam(r, "sessions_id")

	 	attendees := []Attendee{}
		
		result := db.Model(&Attendee{}).
			Select("attendees.*").
			Joins("JOIN attendee_sessions ON attendee_sessions.attendees_id = attendees.id").
			Where("attendee_sessions.sessions_id = ?", sessionsId).			
    		Scan(&attendees)


		w.Header().Set("Content-Type", "application/json")
		if result.Error == nil {
			json.NewEncoder(w).Encode(attendees)
		} else {
			http.Error(w, result.Error.Error(), 500)
			return	
		}
	}	
}


func GetAttendees(db *gorm.DB) http.HandlerFunc {
	return getAll[Attendee](db, stringPtr("attendees.name asc"))
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