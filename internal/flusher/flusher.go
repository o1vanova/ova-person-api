package flusher

import (
	models "github.com/ozonva/ova-person-api/internal/models"
	repo "github.com/ozonva/ova-person-api/internal/repo"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(persons []models.Person) []models.Person
}

type flusher struct {
	chunkSize  int
	personRepo repo.Repo
}

func (f flusher) Flush(persons []interface{}) []interface{} {
	panic("implement me")
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	personRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		personRepo: personRepo,
	}
}
