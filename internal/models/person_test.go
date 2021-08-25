package person

import (
	"fmt"
	"testing"
	"time"
)

func TestPersonString(t *testing.T) {
	now := time.Now()
	const (
		person1Id = iota + 1
		person2Id
	)
	const (
		user1Id = iota + 1
		user2Id
	)
	cases := []struct {
		in  Person
		out string
	}{
		{Person{
			Id:         person1Id,
			UserId:     user1Id,
			FirstName:  "Ivan",
			LastName:   "Ivanov",
			MiddleName: "Ivanovich",
			CreatedAt:  now,
			UpdatedAt:  now,
			DeletedAt:  time.Time{},
		}, fmt.Sprintf("Id: %d, Name: Ivanov Ivan Ivanovich, created at: %s", person1Id, now.String())},
		{Person{
			Id:         person2Id,
			UserId:     user2Id,
			FirstName:  "Petr",
			LastName:   "Petrov",
			MiddleName: "Petrovich",
			CreatedAt:  now,
			UpdatedAt:  now,
			DeletedAt:  time.Time{},
		}, fmt.Sprintf("Id: %d, Name: Petrov Petr Petrovich, created at: %s", person2Id, now.String())},
	}

	for _, c := range cases {
		result := c.in.String()
		if result != c.out {
			t.Errorf("expected: %s, result: %s", c.out, result)
		}
	}
}
