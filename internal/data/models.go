package data

import (
	"database/sql"
	"errors"
)

// custom error message
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Parent model to hold all models
type Models struct {
	Movies MovieModel
	Users  UserModel
}

// method that returns a new Models struct
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
	}
}
