package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	models "github.com/ozonva/ova-person-api/internal/models"
)

// PersonRepo - интерфейс хранилища для сущности Entity
type PersonRepo interface {
	AddPerson(context context.Context, person models.Person) (uint64, error)
	ListPersons(context context.Context, limit, offset uint64) (*[]models.Person, error)
	DescribePerson(context context.Context, personId uint64) (*models.Person, error)
	RemovePerson(context context.Context, personId uint64) (uint64, error)
}

type repository struct {
	db sqlx.DB
}

func (repo repository) AddPerson(context context.Context, person models.Person) (uint64, error) {
	panic("implement me")
}

func (repo repository) ListPersons(context context.Context, limit, offset uint64) (*[]models.Person, error) {
	panic("implement me")
}

func (repo repository) DescribePerson(context context.Context, personId uint64) (*models.Person, error) {
	panic("implement me")
}

func (repo repository) RemovePerson(context context.Context, personId uint64) (uint64, error) {
	panic("implement me")
}

func NewContext(db sqlx.DB) PersonRepo {
	return &repository{db: db}
}
