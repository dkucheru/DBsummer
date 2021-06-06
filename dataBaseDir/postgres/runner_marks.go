package postgres

import (
	"context"
)

type RunnerMarksRepository struct {
	*Repository
}

func (r RunnerMarksRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r RunnerMarksRepository) Get(ctx context.Context) error {
	panic("implement me")
}
