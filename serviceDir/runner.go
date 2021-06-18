package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
	"DBsummer/structs"
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

func (s *runnersService) GetRunnerByID(ctx context.Context, id int) ([]*structs.SheetByID, error) {
	runner, err := s.repository.Runners().GetRunnerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return runner, nil
}

func (s *runnersService) PostRunnerToDataBase(ctx context.Context, runner *pdfReading.ExtractedInformation, teacherId int) error {
	err := s.repository.Runners().PostRunnerToDataBase(ctx, runner, teacherId)
	if err != nil {
		return err
	}

	return nil
}
