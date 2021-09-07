package flusher

import (
	"context"
	"log"

	models "github.com/ozonva/ova-person-api/internal/models"
	repo "github.com/ozonva/ova-person-api/internal/repo"
	utils "github.com/ozonva/ova-person-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(ctx context.Context, persons []models.Person) []models.Person
}

type flusher struct {
	chunkSize  int
	personRepo repo.Repo
}

func (f flusher) Flush(ctx context.Context, persons []models.Person) []models.Person {
	if f.chunkSize < 1 {
		log.Printf("ChunkSize must be positive: %v\n", f.chunkSize)
		return persons
	}
	batches := utils.SplitToBulks(persons, f.chunkSize)

	var unsaved []models.Person
	for _, batch := range batches {
		if personId, err := f.personRepo.AddPersons(ctx, batch); err != nil {
			log.Printf("Error when persons weren't saved: %v\n", personId)
			unsaved = append(unsaved, batch...)
		}
	}

	if len(unsaved) > 0 {
		return unsaved
	}

	return nil
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
