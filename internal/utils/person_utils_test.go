package person_utils

import (
	"fmt"
	"testing"
	"time"

	. "github.com/ozonva/ova-person-api/internal/models"
)

func TestGetMapFromSlice(t *testing.T) {
	now := time.Now()
	const (
		person1Id = iota + 10
		person2Id
		user1Id = iota + 100
		user2Id
	)
	persons := []Person {
		Person{
			Id:         person1Id,
			UserId:     user1Id,
			FirstName:  "Ivan",
			LastName:   "Ivanov",
			MiddleName: "Ivanovich",
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		Person{
			Id:         person2Id,
			UserId:     user2Id,
			FirstName:  "Petr",
			LastName:   "Petrov",
			MiddleName: "Petrovich",
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}
	cases := []struct {
		in  uint64
		out string
	}{
		{person1Id, fmt.Sprintf("Id: %d, Name: Ivanov Ivan Ivanovich, created at: %s", person1Id, now.String())},
		{person2Id, fmt.Sprintf("Id: %d, Name: Petrov Petr Petrovich, created at: %s", person2Id, now.String())},
	}

	personMap, _ := GetMapFromSlice(persons)
	for _, c := range cases {
		result := personMap[c.in].String()
		if result != c.out {
			t.Errorf("expected: %s, result: %s", c.out, result)
		}
	}
}
