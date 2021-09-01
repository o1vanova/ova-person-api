package utils

import (
	"log"

	models "github.com/ozonva/ova-person-api/internal/models"
)

// GetMapFromSlice gets converted slice from structure to map
func GetMapFromSlice(persons []models.Person) (map[uint64]models.Person, error) {
	set := make(map[uint64]models.Person)
	for _, person := range persons {
		set[person.Id] = person
	}
	return set, nil
}

// SplitToBulks gets batches of slice by butchSize
func SplitToBulks(persons []models.Person, butchSize int) [][]models.Person {
	if butchSize < 1 {
		return nil
	}

	batches := make([][]models.Person, 0)
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

// GetPersonsWithoutExcludedPersons This function returns the filtered persons
func GetPersonsWithoutExcludedPersons(persons []models.Person, excludedPersons []models.Person) []models.Person {
	if len(excludedPersons) < 1 {
		return persons
	}

	set, err := GetMapFromSlice(excludedPersons)
	if err != nil {
		log.Println(err)
	}
	newSlice := make([]models.Person, 0)
	for _, person := range persons {
		if _, ok := set[person.Id]; !ok {
			newSlice = append(newSlice, person)
		}
	}
	return newSlice
}
