package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
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

func (s *sheetMarksService) PostSheetMarksToDataBase(ctx context.Context, sheetID int, studentId int, sheetMarks *pdfReading.StudInfoFromPDF) (int, error) {
	id, err := s.repository.SheetMarks().PostSheetMarksToDataBase(ctx, sheetID, studentId, sheetMarks)
	if err != nil {
		return id, err
	}
	return id, nil
}
