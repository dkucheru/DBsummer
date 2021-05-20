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

func (v subjectsService) Create(ctx context.Context, subject structs.Subjects) (structs.Subjects, error) {
	panic("implement me")
}

func (v subjectsService) Get(ctx context.Context, idVid string) (structs.Subjects, error) {
	panic("implement me")
}
