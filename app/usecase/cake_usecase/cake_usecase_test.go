package cake_usecase

import (
	"context"
	"errors"
	"testing"
	"vandyahmad24/maxsol/app/mock/repository"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestCakeUsecase(t *testing.T) {
	mockRepo := repository.MockCakeRepository{}
	ctx := context.Background()
	ctx, closer, sp := util.StartRootSpan(ctx, "Cake Service")
	defer closer.Close()
	defer sp.Finish()
	Convey("Unit Test CakeInteractor", t, func() {
		Convey("Scenario InsertCake", func() {
			Convey("Positive Scenario InsertCake", func() {
				Convey("Succees Insert Cake", func() {
					request := &model.Cake{
						Title:       "TEST",
						Description: "TEST",
						Rating:      1,
						Image:       "TEST",
					}
					mockRepo.On("InsertCake", request).Return(mock.Anything, nil).Once()
					uc := NewCakeUsecase(&mockRepo)
					res, err := uc.CreateCake(ctx, request)
					So(err, ShouldBeNil)
					So(res, ShouldNotBeNil)

				})
			})

			Convey("Negative Scenario InsertCake", func() {
				Convey("Failed If request nil", func() {
					request := &model.Cake{}
					mockRepo.On("InsertCake", request).Return(mock.Anything, nil).Once()
					uc := NewCakeUsecase(&mockRepo)
					res, err := uc.CreateCake(ctx, nil)
					So(err, ShouldNotBeNil)
					So(res, ShouldBeNil)
				})
				Convey("Failed If failed decode ", func() {
					request := &model.Cake{}
					mockRepo.On("InsertCake", request).Return(mock.Anything, nil).Once()
					uc := NewCakeUsecase(&mockRepo)
					res, err := uc.CreateCake(ctx, mock.Anything)
					So(err, ShouldNotBeNil)
					So(res, ShouldBeNil)
				})

			})
		})

		Convey("Scenario GetAllCake", func() {
			Convey("Succees GetAllCake", func() {
				mockRepo.On("GetAll").Return(mock.Anything, nil).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.GetAllCake(ctx)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Failed GetAllCake", func() {
				mockRepo.On("GetAll").Return(nil, errors.New("error")).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.GetAllCake(ctx)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})

		Convey("Scenario GetCake", func() {
			Convey("Succees GetCake", func() {
				mockRepo.On("Get", 1).Return(mock.Anything, nil).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.GetCake(ctx, 1)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Failed GetAllCake", func() {
				mockRepo.On("Get", mock.Anything).Return(nil, errors.New("error")).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.GetCake(ctx, 1)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})

		Convey("Scenario DeleteCake", func() {
			Convey("Succees DeleteCake", func() {
				mockRepo.On("Get", 1).Return(mock.Anything, nil).Once()
				mockRepo.On("Delete", 1).Return(mock.Anything, nil).Once()
				uc := NewCakeUsecase(&mockRepo)
				err := uc.DeleteCake(ctx, 1)
				So(err, ShouldBeNil)
			})
			Convey("Failed DeleteCake When ID Not Found", func() {
				mockRepo.On("Get", mock.Anything).Return(nil, errors.New("error")).Once()
				uc := NewCakeUsecase(&mockRepo)
				err := uc.DeleteCake(ctx, 1)
				So(err, ShouldNotBeNil)
			})
			Convey("Failed DeleteCake In DB", func() {
				mockRepo.On("Get", 1).Return(mock.Anything, nil).Once()
				mockRepo.On("Delete", 1).Return(nil, errors.New("errors")).Once()
				uc := NewCakeUsecase(&mockRepo)
				err := uc.DeleteCake(ctx, 1)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Scenario UpdateCake", func() {
			Convey("Succees UpdateCake", func() {
				request := &model.Cake{}
				mockRepo.On("Get", 1).Return(mock.Anything, nil).Once()
				mockRepo.On("Update", request).Return(mock.Anything, nil).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.UpdateCake(ctx, 1, request)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Failed UpdateCake When ID not found", func() {
				request := &model.Cake{}
				mockRepo.On("Get", 1).Return(nil, errors.New("error")).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.UpdateCake(ctx, 1, request)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("Failed UpdateCake When request nil", func() {
				mockRepo.On("Get", 1).Return(nil, errors.New("error")).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.UpdateCake(ctx, 1, nil)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})

			Convey("Failed UpdateCake When request cannot be decoded", func() {
				mockRepo.On("Get", 1).Return(mock.Anything, nil).Once()
				uc := NewCakeUsecase(&mockRepo)
				res, err := uc.UpdateCake(ctx, 1, mock.Anything)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})

		})

	})

}
