package serviceDir

import (
	"DBsummer/dataBaseDir"
	"context"
)

type runnersService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newRunnersService(service *Service, repository dataBaseDir.Repository) *runnersService {
	subServ := runnersService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *runnersService) Create(ctx context.Context) (int, error) {
	panic("implement me")
}

func (s *runnersService) Get(ctx context.Context) error {
	panic("implement me")
}
