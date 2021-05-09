package dataBaseDir

import (
	"context"
)

type Repository interface {
	Close() error
	BeginTx(ctx context.Context) (Transaction, error)
	TableNew() TableNewInterface
}

type Transaction interface {
	Repository

	Commit() error
	Rollback() error
}

type TableNewInterface interface { //functions to be used for the table 'test'
	Get(ctx context.Context, id int) error
	Create(ctx context.Context) (id int, err error)
}
