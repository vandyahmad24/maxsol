package order_repository

import (
	"vandyahmad24/maxsol/app/model"

	"github.com/opentracing/opentracing-go"
)

type OrderRepository interface {
	InsertOrder(span opentracing.Span, input *model.Order) (interface{}, error)
	GetAll(span opentracing.Span) (interface{}, error)
	Get(span opentracing.Span, id int) (interface{}, error)
	Delete(span opentracing.Span, id int) error
	Update(span opentracing.Span, id int, input *model.Order) (interface{}, error)
}
