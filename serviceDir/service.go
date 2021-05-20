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
}

func NewService(conf *Config) *Service {
	service := &Service{
		repository: conf.Repository,
	}

	service.Subjects = newSubjectsService(service, service.repository)

	return service
}
