package flusher_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-person-api/internal/flusher"
	"github.com/ozonva/ova-person-api/internal/mocks"
	models "github.com/ozonva/ova-person-api/internal/models"
)

var _ = Describe("Test flusher", func() {
	const (
		person1Id = iota + 10
		person2Id
		person3Id
		user1Id = iota + 100
		user2Id
		user3Id
	)
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher flusher.Flusher
		persons     []models.Person
	)

	AfterEach(func() {
		mockCtrl.Finish()
	})

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
		testFlusher = flusher.NewFlusher(3, mockRepo)

		persons = []models.Person{
			models.NewPerson(person1Id, user1Id, "Ivan", "Ivanov", "Ivanovich"),
			models.NewPerson(person2Id, user2Id, "Petr", "Petrov", "Petrovich"),
			models.NewPerson(person3Id, user3Id, "Roman", "Romanov", "Romanovich"),
		}
	})

	Describe("Person storage", func() {
		Context("CRUD", func() {
			It("AddPersons", func() {
				list := persons[:1]
				ctx := context.Background()
				mockRepo.EXPECT().AddPersons(ctx, list).Return(nil)
				Expect(testFlusher.Flush(list)).To(BeNil())
			})
		})
	})
})
