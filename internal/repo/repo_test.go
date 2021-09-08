package repo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	models "github.com/ozonva/ova-person-api/internal/models"
)

var _ = Describe("Repo", func() {
	var (
		personId    uint64
		person      models.Person
	)
	
	BeforeEach(func() {
		person = models.NewPerson(personId, 1, "Ivan", "Ivanov", "Ivanovich")
	})

	Describe("Repo services", func() {
		Context("AddPerson", func() {
			It("should be a novel", func() {
				Expect(person.FirstName).To(Equal("Ivan"))
			})
		})
	})
})
