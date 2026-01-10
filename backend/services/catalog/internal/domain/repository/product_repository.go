package repository

import (
	"context"

	"github.com/kinoshitatakumi/opti/services/catalog/internal/domain/model"
)

type ProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	List(ctx context.Context) ([]*model.Product, error)
	GetByID(ctx context.Context, id model.ProductID) (*model.Product, error)
	// Delete can be added later
}
