package cake_usecase

import (
	"context"
)

type CakeUsecasePort interface {
	CreateCake(ctx context.Context, in interface{}) (interface{}, error)
	GetAllCake(ctx context.Context) (interface{}, error)
	GetCake(ctx context.Context, id int) (interface{}, error)
	DeleteCake(ctx context.Context, id int) error
	UpdateCake(ctx context.Context, id int, in interface{}) (interface{}, error)
}
