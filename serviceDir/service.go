package serviceDir

import (
	db "DBsummer/dataBaseDir"
)

type Config struct {
	Repository db.Repository
}

type Service struct {
	repository db.Repository

	Subjects *subjectsService
	Groups   *groupsService
	Students *studentsService
	Teachers *teachersService
}

func NewService(conf *Config) *Service {
	service := &Service{
		repository: conf.Repository,
	}

	service.Subjects = newSubjectsService(service, service.repository)
	service.Groups = newGroupsService(service, service.repository)
	service.Students = newStudentsService(service, service.repository)
	service.Teachers = newTeachersService(service, service.repository)
	return service
}
