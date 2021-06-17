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

func (s *sheetsService) PostSheetToDataBase(ctx context.Context, sheet *pdfReading.ExtractedInformation, teacherId int, groupId int) (*pdfReading.ExtractedInformation, error) {
	str, err := s.repository.Sheets().PostSheetToDataBase(ctx, sheet, teacherId, groupId)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func (s *sheetsService) GetAVGSheetMark(ctx context.Context, sheetId int) (*float32, error) {
	avgMark, err := s.repository.Sheets().GetAVGSheetMark(ctx, sheetId)
	if err != nil {
		return nil, err
	}

	return avgMark, nil
}

func (s *sheetsService) GetSheetByID(ctx context.Context, id int) ([]*structs.SheetByID, error) {
	sheet, err := s.repository.Sheets().GetSheetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return sheet, nil
}

func (s *sheetsService) DeleteAllData(ctx context.Context) error {
	err := s.repository.Sheets().DeleteAllData(ctx)
	if err != nil {
		return err
	}

	return nil
}
