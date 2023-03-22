package cake_usecase

import (
	"context"
	"errors"
	"vandyahmad24/maxsol/app/domain/repository/cake_repository"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"

	"github.com/mitchellh/mapstructure"
)

type CakeUsecase struct {
	repository cake_repository.CakeRepository
}

func NewCakeUsecase(repository cake_repository.CakeRepository) *CakeUsecase {
	return &CakeUsecase{repository: repository}
}

func (e *CakeUsecase) CreateCake(ctx context.Context, in interface{}) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	util.LogRequest(sp, in)

	var inputCake *model.Cake

	err := mapstructure.Decode(in, &inputCake)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	data, err := e.repository.InsertCake(sp, inputCake)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, data)
	return data, nil
}

func (e *CakeUsecase) GetAllCake(ctx context.Context) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	data, err := e.repository.GetAll(sp)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}
	util.LogResponse(sp, data)

	return data, nil
}

func (e *CakeUsecase) GetCake(ctx context.Context, id int) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	data, err := e.repository.Get(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}
	util.LogResponse(sp, data)

	return data, nil
}

func (e *CakeUsecase) DeleteCake(ctx context.Context, id int) error {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	_, err := e.repository.Get(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return errors.New("Cake Not Fund")
	}

	err = e.repository.Delete(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return err
	}

	return nil
}

func (e *CakeUsecase) UpdateCake(ctx context.Context, id int, in interface{}) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	if in == nil {
		return nil, errors.New("request cannot be nil")
	}

	var inputCake *model.Cake
	err := mapstructure.Decode(in, &inputCake)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	_, err = e.repository.Get(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("Cake Not Fund")
	}

	data, err := e.repository.Update(sp, id, inputCake)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}
	util.LogResponse(sp, data)

	return data, nil
}
