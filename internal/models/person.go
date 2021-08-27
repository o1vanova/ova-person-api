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
	return fmt.Sprintf("Id: %d, Name: %s %s %s, created at: %s",
		person.Id,
		person.LastName,
		person.FirstName,
		person.MiddleName,
		person.CreatedAt.String())
}
