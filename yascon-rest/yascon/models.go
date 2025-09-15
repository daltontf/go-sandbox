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
	Attendees_ID 	int `json:"attendees_id"`
	Sessions_ID 	int	`json:"sessions_id"`
}

type Presentation struct {
	ID       		int    	`json:"id"`	
	Name     		string 	`json:"name"`
	Description 	string 	`json:"description"`
	Speakers_ID 	int 	`json:"speakers_id"`
}

type PresentationWithSpeaker struct {
	ID       		int    	`json:"id"`	
	Name     		string 	`json:"name"`
	Description 	string 	`json:"description"`
	Speaker 		Speaker `json:"speaker" gorm:"embedded;embeddedPrefix:speaker_"` 
}

type Session struct {
	ID       		 int `json:"id"`		
	StartTimeMin	 int `json:"start_time_min"`	
	DurationMins	 int `json:"duration_mins"`	
	Presentations_ID int `json:"presentations_id"`
	Venues_ID 		 int `json:"venues_id"`	
}

type SessionWithPresentationAndVenue struct {
	ID       		 int `json:"id"`		
	StartTimeMin	 int `json:"start_time_min"`	
	DurationMins	 int `json:"duration_mins"`	
	Presentation     Presentation `json:"presentation" gorm:"embedded;embeddedPrefix:presentation_"` 
	Venue 		 	 Venue `json:"venue" gorm:"embedded;embeddedPrefix:venue_"` 
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

type HasID interface { 
	GetID() int
}

func (it Venue) GetID() int { return it.ID }
func (it Speaker) GetID() int { return it.ID }
func (it Attendee) GetID() int { return it.ID }
func (it Presentation) GetID() int { return it.ID }
func (it Session) GetID() int { return it.ID }

func CreateTablesIfNotExists(db *gorm.DB) {
	db.Exec("PRAGMA foreign_keys = ON;")

  	db.Exec(`CREATE TABLE IF NOT EXISTS Attendees (
	  	id INTEGER PRIMARY KEY AUTOINCREMENT,
	  	name TEXT NOT NULL,
	  	password TEXT NOT NULL
	);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS Speakers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		bio TEXT NOT NULL,
		picture BLOB 
	);`)  

	db.Exec(`CREATE TABLE IF NOT EXISTS Presentations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		SPEAKERS_ID INTEGER NOT NULL,
		FOREIGN KEY (SPEAKERS_ID) REFERENCES Speakers(id)
	);`)	 

	db.Exec(`CREATE TABLE IF NOT EXISTS Venues (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		capacity TEXT NOT NULL
	);`)	

	db.Exec(`CREATE TABLE IF NOT EXISTS Sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		start_time_min INTEGER NOT NULL,
		duration_mins INTEGER NOT NULL,
		PRESENTATIONS_ID INTEGER,		
		VENUES_ID INTEGER NOT NULL,
		FOREIGN KEY (PRESENTATIONS_ID) REFERENCES Presentations(id),		
		FOREIGN KEY (VENUES_ID) REFERENCES Venues(id)
	);`) 

	db.Exec(`CREATE TABLE IF NOT EXISTS Attendee_Sessions (
		attendees_id INTEGER NOT NULL,
		sessions_id INTEGER NOT NULL,
		FOREIGN KEY (attendees_id) REFERENCES Attendees(id),
		FOREIGN KEY (sessions_id) REFERENCES Sessions(id)
	);`)
	
	db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_attendee_name ON Attendees(name);	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_attendee_session_attendee ON Attendee_Sessions(attendees_id, sessions_id);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_attendee_session_session ON Attendee_Sessions(sessions_id, attendees_id);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_presentation_name ON Presentations(name);	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_speaker_name ON Speakers(name);	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_venue_name ON Venues(name);	
	`)
}
