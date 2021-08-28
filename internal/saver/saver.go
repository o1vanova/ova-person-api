package saver

import (
	"log"
	"sync"
	"time"

	"github.com/ozonva/ova-person-api/internal/flusher"
	models "github.com/ozonva/ova-person-api/internal/models"
)

type Saver interface {
	Save(person models.Person) // заменить на свою сущность
	Init(tick int)
	Close()
}

type SaveJob struct {
	MtxBuffer *sync.Mutex
	Buffer    []models.Person
	Capacity  uint
	Flusher   flusher.Flusher
	Chanel    chan models.Person
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
) Saver {
	job := SaveJob{
		MtxBuffer: new(sync.Mutex),
		Buffer:    make([]models.Person, capacity),
		Capacity:  capacity,
		Flusher:   flusher,
		Chanel:    make(chan models.Person),
	}
	return &job
}

func (job *SaveJob) GetChanelReadOnly() <-chan models.Person {
	return job.Chanel
}

func (job *SaveJob) Init(tick int) {
	if tick < 1 {
		panic("Time must be positive number")
	}
	ticker := time.NewTicker(time.Duration(tick) * time.Second)

	for {
		select {
		case result := <-job.GetChanelReadOnly():
			log.Printf("\nGot person: %v\n", result.String())
		case <-ticker.C:
			log.Print(".")
		}
	}
}

func (job *SaveJob) Close() {
	close(job.Chanel)
}

func (job *SaveJob) Save(person models.Person) {
}
