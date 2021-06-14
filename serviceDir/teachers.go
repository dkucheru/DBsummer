package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
	"DBsummer/structs"
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

func (s *teachersService) FindTeacher(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error) {
	id, err := s.repository.Teachers().FindTeacher(ctx, sheet)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *teachersService) AddTeacher(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error) {
	id, err := s.repository.Teachers().AddTeacher(ctx, sheet)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *teachersService) GetTeacherPassStatistics(ctx context.Context, passedOrNot string) ([]*structs.TeacherPassStatistics, error) {
	teacherStatistics, err := s.repository.Teachers().GetTeacherPassStatistics(ctx, passedOrNot)
	if err != nil {
		return nil, err
	}

	return teacherStatistics, nil
}
