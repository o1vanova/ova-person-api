package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	models "github.com/ozonva/ova-person-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddPersons(context context.Context, persons []models.Person) (uint64, error)
	ListPersons(context context.Context, limit, offset uint64) ([]models.Person, error)
	DescribePerson(context context.Context, personId uint64) (*models.Person, error)
}

type repository struct {
	db sqlx.DB
}

func (repo repository) AddPersons(context context.Context, persons []models.Person) (uint64, error) {
	panic("implement me")
}

func (repo repository) ListPersons(context context.Context, limit, offset uint64) ([]models.Person, error) {
	panic("implement me")
}

func (repo repository) DescribePerson(context context.Context, personId uint64) (*models.Person, error) {
	panic("implement me")
}

func NewContext(db sqlx.DB) Repo {
	return &repository{db: db}
}
