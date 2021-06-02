package dataBaseDir

import (
	"DBsummer/structs"
	"context"
)

type Repository interface {
	Close() error
	BeginTx(ctx context.Context) (Transaction, error)
	TableNew() TableNewInterface
	Subjects() SubjectsInterface
	Groups() GroupsInterface
}

type Transaction interface {
	Repository

	Commit() error
	Rollback() error
}

type TableNewInterface interface { //functions to be used for the table 'tablen'
	Get(ctx context.Context, id int) error
	Create(ctx context.Context) (id int, err error)
}

type SubjectsInterface interface { //functions to be used for the table 'subjects'
	Get(ctx context.Context, id int) (*structs.Subject, error)
	GetAll(ctx context.Context) ([]*structs.Subject, error)
	Create(ctx context.Context) (id int, err error)
}

type GroupsInterface interface {
}
