package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
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

func (s *runnersService) PostRunnerToDataBase(ctx context.Context, runner *pdfReading.ExtractedInformation, teacherId int) error {
	err := s.repository.Runners().PostRunnerToDataBase(ctx, runner, teacherId)
	if err != nil {
		return err
	}

	return nil
}
