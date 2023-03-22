package user_usecase

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"vandyahmad24/maxsol/app/domain/repository/user_repository"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"
)

type UserUsecase struct {
	repository user_repository.UserRepository
}

func NewUserUsecase(repository user_repository.UserRepository) *UserUsecase {
	return &UserUsecase{repository}
}

func (s *UserUsecase) RegisterUser(ctx context.Context, input model.User) (*model.User, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	util.LogRequest(sp, input)
	//find unique username
	oldUser, err := s.repository.FindByName(sp, input.Name)
	if err != nil {
		return nil, err
	}

	if oldUser.Id != 0 {
		return nil, errors.New("username already registered")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	reqUser := model.User{
		Name:     input.Name,
		Password: string(passwordHash),
	}

	user, err := s.repository.Store(sp, reqUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserUsecase) LoginUsernamePass(ctx context.Context, input model.User) (*model.User, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	util.LogRequest(sp, input)

	username := input.Name
	password := input.Password
	user, err := s.repository.FindByName(sp, username)
	if err != nil {
		return user, errors.New("Username/Password wrong")
	}

	if user.Name == "" {
		return user, errors.New("Username/Password wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Username/Password wrong")
	}

	return user, nil
}
