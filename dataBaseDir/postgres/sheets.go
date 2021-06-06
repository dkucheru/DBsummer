package postgres

import "context"

type SheetsRepository struct {
	*Repository
}

func (r SheetsRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r SheetsRepository) Get(ctx context.Context) error {
	panic("implement me")
}
