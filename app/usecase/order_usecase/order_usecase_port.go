package order_usecase

import (
	"context"
)

type OrderUsecasePort interface {
	CreateOrder(ctx context.Context, in interface{}) (interface{}, error)
	GetAllOrder(ctx context.Context) (interface{}, error)
	GetOrder(ctx context.Context, id int) (interface{}, error)
	DeleteOrder(ctx context.Context, id int) error
	UpdateOrder(ctx context.Context, id int, in interface{}) (interface{}, error)
	CreateOrderBulk(ctx context.Context, in interface{}) (interface{}, error)
}
