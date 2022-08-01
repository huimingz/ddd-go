package persistence

import (
	"context"
	"github.com/huimingz/ddd-go/domain/aggregate"
	"github.com/huimingz/ddd-go/domain/port/repository"
)

var _ repository.AuthorRepositoryInterface = &AuthorRepository{}

type AuthorRepository struct{}

func (r *AuthorRepository) GetByID(ctx context.Context, id string) *aggregate.Author {
	// TODO implement me
	panic("implement me")
}

func (r *AuthorRepository) Save(ctx context.Context, author *aggregate.Author) error {
	// TODO implement me
	panic("implement me")
}
