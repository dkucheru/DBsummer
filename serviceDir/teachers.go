package serviceDir

import (
	"DBsummer/dataBaseDir"
	"context"
)

type teachersService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newTeachersService(service *Service, repository dataBaseDir.Repository) *teachersService {
	subServ := teachersService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *teachersService) Create(ctx context.Context) (int, error) {
	panic("implement me")
}

func (s *teachersService) Get(ctx context.Context) error {
	panic("implement me")
}
