package repository

import (
	"context"
	"github.com/huimingz/ddd-go/domain/aggregate"
)

type AuthorRepositoryInterface interface {
	GetByID(ctx context.Context, id string) *aggregate.Author
	Save(ctx context.Context, author *aggregate.Author) error
}
