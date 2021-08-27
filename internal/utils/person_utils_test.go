package utils

import (
	"testing"

	models "github.com/ozonva/ova-person-api/internal/models"
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
		in  [][]models.Person
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

func getTestPersons() []models.Person {
	const (
		person1Id = iota + 10
		person2Id
		person3Id
		user1Id = iota + 100
		user2Id
		user3Id
	)
	return []models.Person{
		models.NewPerson(person1Id, user1Id, "Ivan", "Ivanov", "Ivanovich"),
		models.NewPerson(person2Id, user2Id, "Petr", "Petrov", "Petrovich"),
		models.NewPerson(person3Id, user3Id, "Roman", "Romanov", "Romanovich"),
	}
}
