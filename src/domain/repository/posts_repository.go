package repository

import (
	"context"
	"github.com/nejtr0n/api_example/domain/model"
)

type ProductsRepository interface {
	Find(ctx context.Context, limit, offset int64, sortBy string, sortOrder int) ([]*model.Product, error)
	SaveAll(ctx context.Context, pipe chan *model.Product) (int64, error)
}
