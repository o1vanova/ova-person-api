package saver

import (
	"log"
	"sync"
	"time"

	"github.com/ozonva/ova-person-api/internal/flusher"
	models "github.com/ozonva/ova-person-api/internal/models"
	utils "github.com/ozonva/ova-person-api/internal/utils"
)

type Saver interface {
	Save(person models.Person) // заменить на свою сущность
	Close()
}

type saveJob struct {
	mtxBuffer *sync.Mutex
	buffer    []models.Person
	capacity  uint
	flusher   flusher.Flusher
	channel   chan models.Person
	ticker    time.Ticker
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	tick uint,
) Saver {
	if tick == 0 {
		panic("Time must be more than zero")
	}

	ticker := time.NewTicker(time.Duration(tick) * time.Second)
	job := saveJob{
		mtxBuffer: new(sync.Mutex),
		buffer:    make([]models.Person, capacity),
		capacity:  capacity,
		flusher:   flusher,
		channel:   make(chan models.Person),
		ticker:    *ticker,
	}

	go job.init()
	return &job
}

func (job *saveJob) GetChanelReadOnly() <-chan models.Person {
	return job.channel
}

func (job *saveJob) Close() {
	log.Println("User wants to close")
	job.ticker.Stop()
	close(job.channel)
	job.saveInStorage()
}

func (job *saveJob) Save(person models.Person) {
	log.Println("User wants to save person: ", person.String())
	if job.isBuffLocked() {
		job.saveInStorage()
	}
	job.saveInCash(person)
}

func (job *saveJob) init() {
	for {
		select {
		case result := <-job.GetChanelReadOnly():
			job.Save(result)
		case <-job.ticker.C:
			job.saveInStorage()
		}
	}
}

func (job *saveJob) saveInCash(person models.Person) {
	job.mtxBuffer.Lock()
	defer job.mtxBuffer.Unlock()

	if !job.isBuffLocked() {
		job.buffer = append(job.buffer, person)
	}
}

func (job *saveJob) saveInStorage() {
	job.mtxBuffer.Lock()
	defer job.mtxBuffer.Unlock()

	if len(job.buffer) == 0 {
		return
	}

	unsavedPersons := job.flusher.Flush(job.buffer)
	if len(unsavedPersons) == 0 {
		job.buffer = nil
	} else {
		result := utils.GetPersonsWithoutExcludedPersons(job.buffer, unsavedPersons)
		job.buffer = result
	}
}

func (job *saveJob) isBuffLocked() bool {
	result := cap(job.buffer) == len(job.buffer)
	if result {
		log.Println("Buffer is full! Clear, please")
	}
	return result
}
