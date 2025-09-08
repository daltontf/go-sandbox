package yascon

import (
  	"gorm.io/gorm"
)

type Attendee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AttendeeSession struct {
	ID       		int	`json:"id"`		
	Attendee_ID 	int `json:"attendee_id"`
	Session_ID 		int	`json:"session_id"`
}

type Presentation struct {
	ID       	int    	`json:"id"`	
	Name     	string 	`json:"name"`
	Description string 	`json:"description"`
}

type Session struct {
	ID       		int	`json:"id"`		
	SessionTime_ID 	int `json:"session_time_id"`
	Presentation_ID int `json:"presentation_id"`
	Speaker_ID 		int `json:"speaker_id"`
	Venue_ID 		int `json:"venue_id"`	
}

type SessionTime struct {
	ID       		int	`json:"id"`
	StartTime		int `json:"start_time"`	
	EndTime			int `json:"end_time"`	
}

type Speaker struct {
	ID       		int		`json:"id"`
	Name     		string 	`json:"name"`
	Bio				string 	`json:"bio"`
	Picture			[]byte  `json:"picture"`
}

type Venue struct {
	ID       		int		`json:"id"`
	Name     		string 	`json:"name"`	
	Capacity		int  	`json:"capacity"`	
}

func CreateTablesIfNotExists(db *gorm.DB) {
  	db.Exec(`CREATE TABLE IF NOT EXISTS Attendees (
	  	id INTEGER PRIMARY KEY AUTOINCREMENT,
	  	name TEXT NOT NULL,
	  	password TEXT NOT NULL
	);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS AttendeeSessions (
		attendee_id INTEGER NOT NULL,
		session_id INTEGER NOT NULL		
	);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS Presentations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL
	);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS Sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		SESSION_TIME_ID INTEGER,
		PRESENTATION_TIME_ID INTEGER,
		SPEAKER_ID INTEGER NOT NULL,
		VENUE_ID INTEGER NOT NULL
	);`)  

	db.Exec(`CREATE TABLE IF NOT EXISTS SessionTimes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		start_time INTEGER NOT NULL,
		end_time INTEGER NOT NULL
	);`)  

	db.Exec(`CREATE TABLE IF NOT EXISTS Speakers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		bio TEXT NOT NULL,
		picture BLOB 
	);`)  

	db.Exec(`CREATE TABLE IF NOT EXISTS Venues (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		capacity TEXT NOT NULL
	);`)	
	
	db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_attendee_name ON Attendees(name);	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_attendee_session_attendee ON AttendeeSessions(attendee_id, session_id);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_attendee_session_session ON AttendeeSessions(session_id, attendee_id);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_presentation_name ON Presentations(name);	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_speaker_name ON Speakers(name);	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_venue_name ON Venues(name);	
	`)
}
