package person

import (
	"fmt"
	"testing"
)

func TestPersonString(t *testing.T) {
	const (
		person1Id = iota + 1
		person2Id
		user1Id = iota + 1
		user2Id
	)
	cases := []struct {
		in  Person
		out string
	}{
		{NewPerson(person1Id, user1Id, "Ivan", "Ivanov", "Ivanovich"),
			fmt.Sprintf("Id: %d, Name: Ivanov Ivan Ivanovich", person1Id)},

		{NewPerson(person2Id, user2Id, "Petr", "Petrov", "Petrovich"),
			fmt.Sprintf("Id: %d, Name: Petrov Petr Petrovich", person2Id)},
	}

	for _, c := range cases {
		result := c.in.String()
		if result != c.out {
			t.Errorf("expected: %s, result: %s", c.out, result)
		}
	}
}
