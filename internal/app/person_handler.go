package person_service

import (
	"context"

	desc "github.com/ozonva/ova-person-api/pkg"
	"github.com/rs/zerolog/log"
)

func (a *PersonApi) CreatePersonV1(ctx context.Context, req *desc.CreatePersonRequest) (*desc.CreatePersonResponse, error) {
	log.Info().Msgf("Create person: %s", req.String())
	return &desc.CreatePersonResponse{}, nil
}

func (a *PersonApi) DescribePersonV1(ctx context.Context, req *desc.DescribePersonRequest) (*desc.DescribePersonResponse, error) {
	log.Info().Msgf("Describe person with id: %v", req.PersonId)
	return &desc.DescribePersonResponse{}, nil
}

func (a *PersonApi) ListPersonsPersonV1(ctx context.Context, req *desc.ListPersonsRequest) (*desc.ListPersonsResponse, error) {
	log.Info().Msgf("Get persons %v", req.String())
	return &desc.ListPersonsResponse{}, nil
}

func (a *PersonApi) RemovePersonV1(ctx context.Context, req *desc.RemovePersonRequest) (*desc.RemovePersonResponse, error) {
	log.Info().Msgf("Remove person with id: %v", req.PersonId)
	return &desc.RemovePersonResponse{}, nil
}
