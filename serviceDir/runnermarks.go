package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
	"DBsummer/structs"
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

func (s *runnerMarksService) PostRunnerMarksToDataBase(ctx context.Context, sheetMarkID int, runnerID int, runnerMarks *pdfReading.StudInfoFromPDF) error {
	err := s.repository.RunnerMarks().PostRunnerMarksToDataBase(ctx, sheetMarkID, runnerID, runnerMarks)
	if err != nil {
		return err
	}
	return nil
}

func (s *runnerMarksService) GetRatingStudentWithRunners(ctx context.Context, sem string, ed_y string) ([]*structs.RatingWithRunners, error) {
	ratings, err := s.repository.RunnerMarks().GetRatingStudentWithRunners(ctx, sem, ed_y)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}
