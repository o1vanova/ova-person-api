package app

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	models "github.com/ozonva/ova-person-api/internal/models"
	desc "github.com/ozonva/ova-person-api/pkg"
	"github.com/rs/zerolog/log"
)

func (api *PersonApi) CreatePersonV1(ctx context.Context, req *desc.CreatePersonRequest) (*desc.CreatePersonResponse, error) {
	log.Info().Msgf("Create person: %s", req.String())
	personId, err := api.personRepo.AddPerson(ctx, models.NewPerson(req.PersonId, req.UserId, req.FirstName, req.LastName, req.MiddleName))
	if err != nil {
		log.Error().Msgf("CreatePersonV1 error: %s", err)
		return nil, err
	}
	return &desc.CreatePersonResponse{PersonId: personId}, nil
}

func (api *PersonApi) DescribePersonV1(ctx context.Context, req *desc.DescribePersonRequest) (*desc.DescribePersonResponse, error) {
	log.Info().Msgf("Describe person with id: %v", req.PersonId)
	person, err := api.personRepo.DescribePerson(ctx, req.PersonId)
	if err != nil {
		log.Error().Msgf("DescribePersonV1 error: %s", err)
		return nil, err
	}

	return &desc.DescribePersonResponse{Person: mapPersonToDto(person)}, nil
}

func (api *PersonApi) ListPersonsPersonV1(ctx context.Context, req *desc.ListPersonsRequest) (*desc.ListPersonsResponse, error) {
	log.Info().Msgf("Get persons %v", req.String())
	persons, err := api.personRepo.ListPersons(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Msgf("ListPersonsPersonV1 error: %s", err)
		return nil, err
	}
	var result []*desc.CreatePersonRequest
	for key, person := range persons {
		result[key] = mapPersonToDto(&person)
	}
	return &desc.ListPersonsResponse{Persons: result}, nil
}

func (api *PersonApi) RemovePersonV1(ctx context.Context, req *desc.RemovePersonRequest) (*desc.RemovePersonResponse, error) {
	log.Info().Msgf("Remove person with id: %v", req.PersonId)
	personId, err := api.personRepo.RemovePerson(ctx, req.PersonId)
	if err != nil {
		log.Error().Msgf("RemovePersonV1 error: %s", err)
		return nil, err
	}
	return &desc.RemovePersonResponse{PersonId: personId}, nil
}

func mapPersonToDto(person *models.Person) *desc.CreatePersonRequest {
	result := desc.CreatePersonRequest{
		PersonId:   uint64(person.Id),
		UserId:     person.UserId,
		FirstName:  person.FirstName,
		LastName:   person.LastName,
		MiddleName: person.MiddleName,
		CreatedAt:  timestamppb.New(person.CreatedAt),
		UpdatedAt:  nil,
		DeletedAt:  nil,
	}
	return &result
}
