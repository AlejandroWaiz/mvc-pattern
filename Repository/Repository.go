package repository

import (
	"database/sql"

	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {

	repository := Repository{
		db: db,
	}

	return &repository
}

type IRepository interface {
	CreatePokemonRepo([]models.Pokemon) error
	GetAllPokemonsRepo() ([]models.PokemonWithID, error)
}
