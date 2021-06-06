package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/structs"
	"context"
)

type studentsService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newStudentsService(service *Service, repository dataBaseDir.Repository) *studentsService {
	subServ := studentsService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *studentsService) Create(ctx context.Context, subject structs.Student) (structs.Student, error) {
	panic("implement me")
}

func (s *studentsService) Get(ctx context.Context, idVid int) (*structs.Student, error) {
	panic("implement me")
}

func (s *studentsService) GetAllStudInfo(ctx context.Context) ([]*structs.AllStudInfo, error) {
	subj, err := s.repository.Students().GetAllStudInfo(ctx)
	if err != nil {
		return nil, err
	}

	return subj, nil
}

func (s *studentsService) GetAllBorjniki(ctx context.Context) ([]*structs.AllStudInfo, error) {
	subj, err := s.repository.Students().GetAllBorjniki(ctx)
	if err != nil {
		return nil, err
	}

	return subj, nil
}
