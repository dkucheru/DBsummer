package dataBaseDir

import (
	"DBsummer/configDir"
	"DBsummer/dataBaseDir/postgres"
	"context"
	"log"
)

type DatabaseRepository struct {
	DBRepository *postgres.Repository
	Config       configDir.DBConfig
	Log          *log.Logger
}

func NewDatabaseRepository(log *log.Logger, config configDir.DBConfig) (*DatabaseRepository, error) {
	repository := &DatabaseRepository{
		Log:    log,
		Config: config,
	}

	var err error
	repository.DBRepository, err = postgres.New(log, config)
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (d *DatabaseRepository) Close() error {
	return d.DBRepository.Close()
}

func (d *DatabaseRepository) TableNew() TableNewInterface {
	return d.DBRepository.TableNew()
}

func (d *DatabaseRepository) Subjects() SubjectsInterface {
	return d.DBRepository.Subjects()
}

func (d *DatabaseRepository) Groups() GroupsInterface {
	return d.DBRepository.Groups()
}

func (d *DatabaseRepository) BeginTx(ctx context.Context) (Transaction, error) {
	tx, err := d.DBRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	r := *d

	r.DBRepository = tx

	return &DBTx{
		DatabaseRepository: &r,
	}, nil
}

type DBTx struct {
	*DatabaseRepository
}

func (d *DBTx) Commit() error {
	return d.DBRepository.Commit()
}

func (d *DBTx) Rollback() error {
	return d.DBRepository.Rollback()
}
