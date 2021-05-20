package postgres

import (
	"context"
	"math/rand"
)

type SubjectsRepository struct {
	*Repository
}

func (r SubjectsRepository) Create(ctx context.Context) (id int, err error) {
	query := r.db.Rebind(`
		INSERT into tablen(id_t) VALUES(?);
	`)

	id = int(rand.Uint32())

	_, err = r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r SubjectsRepository) Get(ctx context.Context, id int) error {
	panic("implement me")
}
