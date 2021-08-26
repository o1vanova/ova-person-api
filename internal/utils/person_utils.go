package person_utils

import (
	. "github.com/ozonva/ova-person-api/internal/models"
)

// GetMapFromSlice Get converted slice from structure to map,
func GetMapFromSlice(persons []Person) (map[uint64]Person, error) {
	set := make(map[uint64]Person)
	for _, person := range persons {
		set[person.Id] = person
	}
	return set, nil
}

// SplitToBulks Get batches of slice by butchSize,
func SplitToBulks(persons []Person, butchSize uint) [][]Person {
	if butchSize < 1 {
		return nil
	}

	len := uint(len(persons))
	count := len / butchSize
	batches := make([][]Person, 0, butchSize)

	for i := uint(0); i < butchSize; i++ {
		from := i * count
		batches = append(batches, persons[from:from+count])
	}

	if len%butchSize > 0 {
		batches[butchSize-1] = append(batches[butchSize-1], persons[butchSize*count:]...)
	}
	return batches
}
