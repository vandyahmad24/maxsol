package user_repository

import (
	"vandyahmad24/maxsol/app/model"

	"github.com/opentracing/opentracing-go"
)

type UserRepository interface {
	FindByName(span opentracing.Span, username string) (*model.User, error)
	Store(span opentracing.Span, user model.User) (*model.User, error)
}
