package website

import "context"

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, website Website) (*Website, error)
	All(ctx context.Context) ([]Website, error)
	GetByNamd(ctx context.Context, name string) (*Website, error)
	Update(ctx context.Context, id int64, update Website) (*Website, error)
	Delete(ctx context.Context, id int64) error
}
