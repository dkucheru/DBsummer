package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
	"DBsummer/structs"
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

func (s *sheetsService) GetSheetFromParameters(ctx context.Context, fn string, ln string, mn string, subj string, gr string, year string) ([]*structs.SheetByQuery, error) {
	struc, err := s.repository.Sheets().GetSheetFromParameters(ctx, fn, ln, mn, subj, gr, year)
	if err != nil {
		return nil, err
	}

	return struc, nil
}

func (s *sheetsService) PostSheetToDataBase(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*pdfReading.ExtractedInformation, error) {
	str, err := s.repository.Sheets().PostSheetToDataBase(ctx, sheet)
	if err != nil {
		return nil, err
	}

	return str, nil
}
