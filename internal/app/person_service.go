package person_service

import (
	repo "github.com/ozonva/ova-person-api/internal/repo"
	desc "github.com/ozonva/ova-person-api/pkg"
)

type PersonApi struct {
	desc.UnimplementedPersonApiServiceServer
	personRepo repo.PersonRepo
}

func NewPersonApi(personRepo repo.PersonRepo) desc.PersonApiServiceServer {
	return &PersonApi{
		personRepo: personRepo,
	}
}
