package person_utils

import (
	"testing"
	"time"

	"github.com/ozonva/ova-person-api/internal/models/person"
)

func TestGetMapFromSlice(t *testing.T) {
	now := time.Now()
	cases := []struct {
		in  Person
		out int
	}{
		{Person{
			Id:         17,
			UserId:     1,
			FirstName:  "Ivan",
			LastName:   "Ivanov",
			MiddleName: "Ivanovich",
			CreatedAt:  now,
			UpdatedAt:  now,
		}, 17},
	}

	for _, c := range cases {
		result := GetMapFromSlice(c).in.String()
		if result[c.out] != c.out {
			t.Errorf("expected: %s, result: %s", c.out, result[c.out])
		}
	}
}
