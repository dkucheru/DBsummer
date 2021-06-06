package postgres

import (
	"context"
)

type TeachersRepository struct {
	*Repository
}

func (r TeachersRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r TeachersRepository) Get(ctx context.Context) error {
	panic("implement me")
}
