package database

import "database/sql"


type Models struct {
	Users UsersModel
	Events EventsModel
	Attendees AttendeesModel
}

func NewModels(db *sql.DB) Models {
	return  Models{
		Users: UsersModel{DB: db},
		Events: EventsModel{DB: db},
		Attendees: AttendeesModel{DB: db},
	}
}