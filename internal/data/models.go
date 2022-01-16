package data

import (
	"database/sql"
	"errors"
)

//this error will be return from Get() method when movie doesn't exist in db
var (
	ErrRecordNotFound = errors.New("record not found")
)

//Model wraps MovieModel
type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
    return Models{
        Movies: MovieModel{DB: db},
    }
}