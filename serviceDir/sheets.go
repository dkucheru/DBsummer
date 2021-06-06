package serviceDir

import (
	"DBsummer/dataBaseDir"
	"context"
)

type sheetsService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newSheetsService(service *Service, repository dataBaseDir.Repository) *sheetsService {
	subServ := sheetsService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *sheetsService) Create(ctx context.Context) (int, error) {
	panic("implement me")
}

func (s *sheetsService) Get(ctx context.Context) error {
	panic("implement me")
}
