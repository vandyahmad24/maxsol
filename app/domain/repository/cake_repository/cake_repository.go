package cake_repository

import (
	"vandyahmad24/maxsol/app/model"

	"github.com/opentracing/opentracing-go"
)

type CakeRepository interface {
	InsertCake(span opentracing.Span, input *model.Cake) (interface{}, error)
	GetAll(span opentracing.Span) (interface{}, error)
	Get(span opentracing.Span, id int) (interface{}, error)
	Delete(span opentracing.Span, id int) error
	Update(span opentracing.Span, id int, input *model.Cake) (interface{}, error)
}
