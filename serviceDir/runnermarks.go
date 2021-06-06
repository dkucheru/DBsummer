package serviceDir

import (
	"DBsummer/dataBaseDir"
	"context"
)

type runnerMarksService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newrunnerMarksService(service *Service, repository dataBaseDir.Repository) *runnerMarksService {
	subServ := runnerMarksService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *runnerMarksService) Create(ctx context.Context) (int, error) {
	panic("implement me")
}

func (s *runnerMarksService) Get(ctx context.Context) error {
	panic("implement me")
}
