package serviceDir

import (
	"DBsummer/dataBaseDir"
	"context"
)

type sheetMarksService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newSheetMarksService(service *Service, repository dataBaseDir.Repository) *sheetMarksService {
	subServ := sheetMarksService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *sheetMarksService) Create(ctx context.Context) (int, error) {
	panic("implement me")
}

func (s *sheetMarksService) Get(ctx context.Context) error {
	panic("implement me")
}
