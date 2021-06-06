package postgres

import "context"

type SheetMarksRepository struct {
	*Repository
}

func (r SheetMarksRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r SheetMarksRepository) Get(ctx context.Context) error {
	panic("implement me")
}
