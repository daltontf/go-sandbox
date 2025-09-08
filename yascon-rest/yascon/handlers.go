package yascon

import (
	"github.com/go-chi/chi/v5"
	"context"
	"encoding/json"
	"net/http"
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

func create[T any](db *gorm.DB) http.HandlerFunc {
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

func CreateSpeaker(db *gorm.DB) http.HandlerFunc {
	return create[Speaker](db)
}

func DeleteSpeaker(db *gorm.DB) http.HandlerFunc {
	return deleteById[Speaker](db)
}

func UpdateSpeaker(db *gorm.DB) http.HandlerFunc {
	return updateById[Speaker](db)
}



func GetSessionTimes(db *gorm.DB) http.HandlerFunc {
	return getAll[SessionTime](db)
}

func CreateSessionTime(db *gorm.DB) http.HandlerFunc {
	return create[SessionTime](db)
}

func DeleteSessionTime(db *gorm.DB) http.HandlerFunc {
	return deleteById[SessionTime](db)
}

func UpdateSessionTime(db *gorm.DB) http.HandlerFunc {
	return updateById[SessionTime](db)
}


func GetSessions(db *gorm.DB) http.HandlerFunc {
	return getAll[Session](db)
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
	return getAll[Presentation](db)
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



func GetAttendeeSessions(db *gorm.DB) http.HandlerFunc {
	return getAll[AttendeeSession](db)
}

func CreateAttendeeSession(db *gorm.DB) http.HandlerFunc {
	return create[AttendeeSession](db)
}

func DeleteAttendeeSession(db *gorm.DB) http.HandlerFunc {
	return deleteById[AttendeeSession](db)
}

func UpdateAttendeeSession(db *gorm.DB) http.HandlerFunc {
	return updateById[AttendeeSession](db)
}


func GetAttendees(db *gorm.DB) http.HandlerFunc {
	return getAll[Attendee](db)
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