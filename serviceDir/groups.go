package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/structs"
	"context"
)

type groupsService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newGroupsService(service *Service, repository dataBaseDir.Repository) *groupsService {
	subServ := groupsService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *groupsService) Create(ctx context.Context, subject structs.Group) (structs.Group, error) {
	panic("implement me")
}

func (s *groupsService) Get(ctx context.Context, idVid int) (*structs.Group, error) {
	panic("implement me")
}
