package serviceDir

import (
	"awesomeProject/db"
)

type Config struct {
	Repository db.Repository
}

type Service struct {
	Repository db.Repository
}

func NewService(conf *Config) *Service {
	service := &Service{
		Repository: conf.Repository,
	}
	return service
}
