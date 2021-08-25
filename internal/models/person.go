package person

import "time"

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
