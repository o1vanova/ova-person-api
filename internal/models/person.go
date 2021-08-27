package person

import (
	"fmt"
	"time"
)

type Person struct {
	Id         uint64
	UserId     uint64
	FirstName  string
	LastName   string
	MiddleName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func (person Person) String() string {
	return fmt.Sprintf("Id: %d, Name: %s %s %s",
		person.Id,
		person.LastName,
		person.FirstName,
		person.MiddleName)
}

func NewPerson(personId uint64, userId uint64, firstName string, lastName string, middleName string) Person {
	now := time.Now()
	return Person{
		Id:         personId,
		UserId:     userId,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
