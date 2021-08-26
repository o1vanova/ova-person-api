package utils

import (
	"testing"
	"time"

	. "github.com/ozonva/ova-person-api/internal/models"
)

func TestGetMapFromSlice(t *testing.T) {
	type Case struct {
		in  uint64
		out string
	}
	var cases []Case
	persons := getTestPersons()

	for _, person := range persons {
		cases = append(cases, Case{person.Id, person.String()})
	}

	personMap, _ := GetMapFromSlice(persons)
	for _, c := range cases {
		result := personMap[c.in].String()
		if result != c.out {
			t.Errorf("expected: %s, result: %s", c.out, result)
		}
	}
}

func TestSplitToBulks(t *testing.T) {
	persons := getTestPersons()
	cases := []struct {
		in  [][]Person
		out int
	}{
		{SplitToBulks(persons, 2), 2},
		{SplitToBulks(persons, 3), 3},
	}

	for _, c := range cases {
		result := len(c.in)
		if result != c.out {
			t.Errorf("expected: %d, result: %d", c.out, result)
		}
	}
}

func getTestPersons() []Person {
	now := time.Now()
	const (
		person1Id = iota + 10
		person2Id
		person3Id
		user1Id = iota + 100
		user2Id
		user3Id
	)
	return []Person {
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
		Person{
			Id:         person3Id,
			UserId:     user3Id,
			FirstName:  "Roman",
			LastName:   "Romanov",
			MiddleName: "Romanich",
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}
}
