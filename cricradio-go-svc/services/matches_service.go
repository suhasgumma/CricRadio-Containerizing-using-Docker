package services

import (
	"cricradio-go-svc/domain/matches"
	"cricradio-go-svc/utils/errors"
)

var (
	MatchesService MatchesServiceInterface = &matchesService{}
)

type matchesService struct {
}

type MatchesServiceInterface interface {
	List() (matches.Matches, *errors.RestErr)
}

func (m matchesService) List() (matches.Matches, *errors.RestErr) {
	mt := &matches.Match{}
	return mt.ListMatches()
}
