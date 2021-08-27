package utils

import (
	. "github.com/ozonva/ova-person-api/internal/models"
)

// GetMapFromSlice gets converted slice from structure to map
func GetMapFromSlice(persons []Person) (map[uint64]Person, error) {
	set := make(map[uint64]Person)
	for _, person := range persons {
		set[person.Id] = person
	}
	return set, nil
}

// SplitToBulks gets batches of slice by butchSize
func SplitToBulks(persons []Person, butchSize int) [][]Person {
	if butchSize < 1 {
		return nil
	}

	batches := make([][]Person, 0)
	chunkSize := (len(persons) + butchSize - 1) / butchSize

	for from := 0; from < len(persons); from += chunkSize {
		to := from + chunkSize

		if to > len(persons) {
			to = len(persons)
		}

		batches = append(batches, persons[from:to])
	}
	return batches
}
