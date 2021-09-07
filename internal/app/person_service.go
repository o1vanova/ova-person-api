package person_service

import (
	desc "github.com/ozonva/ova-person-api/pkg"
)

type PersonApi struct {
	desc.UnimplementedPersonApiServiceServer
}

func NewPersonApi() desc.PersonApiServiceServer {
	return &PersonApi{}
}
