package repo

import (
	models "github.com/ozonva/ova-person-api/internal/models"
)

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddPersons(persons []models.Person) error
	ListPersons(limit, offset uint64) ([]models.Person, error)
	DescribePerson(personId uint64) (*models.Person, error)
}
