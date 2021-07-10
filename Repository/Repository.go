package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {

	repository := Repository{
		db: db,
	}

	return &repository
}

type IRepository interface {
	CreatePokemon()
}
