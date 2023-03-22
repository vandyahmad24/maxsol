package user_usecase

import (
	"context"
	"vandyahmad24/maxsol/app/model"
)

type UserUsecasePort interface {
	RegisterUser(ctx context.Context, input model.User) (*model.User, error)
	LoginUsernamePass(ctx context.Context, input model.User) (*model.User, error)
}
