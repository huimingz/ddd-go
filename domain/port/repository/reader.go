package repository

import (
	"context"
	"github.com/huimingz/ddd-go/domain/aggregate"
)

type ReaderRepositoryInterface interface {
	GetByID(ctx context.Context, id string) *aggregate.Reader
	Save(ctx context.Context, reader *aggregate.Reader) error
}
