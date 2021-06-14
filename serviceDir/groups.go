package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
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

func (s *groupsService) FindGroup(ctx context.Context, sheet *pdfReading.ExtractedInformation, subjectId int) (*int, error) {
	id, err := s.repository.Groups().FindGroup(ctx, sheet, subjectId)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *groupsService) AddGroup(ctx context.Context, sheet *pdfReading.ExtractedInformation, subjectId int) (*int, error) {
	id, err := s.repository.Groups().AddGroup(ctx, sheet, subjectId)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *groupsService) FindGroupsOfScientist(ctx context.Context, scientificDegree string) ([]*structs.GroupOfScientist, error) {
	groups, err := s.repository.Groups().FindGroupsOfScientist(ctx, scientificDegree)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
