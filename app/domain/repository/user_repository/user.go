package user_repository

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Store(span opentracing.Span, user model.User) (*model.User, error) {
	sp := util.CreateSubChildSpan(span, "Store")
	defer sp.Finish()

	err := r.db.Create(&user).Error
	if err != nil {
		util.LogError(sp, err)
		return &user, err
	}
	util.LogResponse(sp, user)
	return &user, nil
}

func (r *userRepository) FindByName(span opentracing.Span, username string) (*model.User, error) {
	sp := util.CreateSubChildSpan(span, "FindByName")
	defer sp.Finish()

	var user model.User
	err := r.db.Where("name = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	util.LogResponse(sp, user)

	return &user, nil

}
