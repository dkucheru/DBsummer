package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/structs"
	"context"
)

type subjectsService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newSubjectsService(service *Service, repository dataBaseDir.Repository) *subjectsService {
	subServ := subjectsService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *subjectsService) Create(ctx context.Context, subject structs.Subject) (structs.Subject, error) {
	panic("implement me")
}

//func (s *subjectsService) Get(ctx context.Context, idVid int) (*structs.Subject, error) {
func (s *subjectsService) Get(ctx context.Context, idVid int) (*structs.Subject, error) {
	subj, err := s.repository.Subjects().Get(ctx, idVid)
	if err != nil {
		return nil, err
	}

	return subj, nil
}

func (s *subjectsService) GetAll(ctx context.Context) ([]*structs.Subject, error) {
	subj, err := s.repository.Subjects().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return subj, nil
}
