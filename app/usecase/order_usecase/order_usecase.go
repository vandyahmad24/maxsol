package order_usecase

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"vandyahmad24/maxsol/app/domain/entity"
	"vandyahmad24/maxsol/app/domain/repository/cake_repository"
	"vandyahmad24/maxsol/app/domain/repository/order_repository"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"

	"github.com/mitchellh/mapstructure"
)

type OrderUsecase struct {
	repository     order_repository.OrderRepository
	cakeRepository cake_repository.CakeRepository
}

func NewOrderUsecase(repository order_repository.OrderRepository, cakeRepository cake_repository.CakeRepository) *OrderUsecase {
	return &OrderUsecase{repository: repository, cakeRepository: cakeRepository}
}

func (e *OrderUsecase) CreateOrder(ctx context.Context, in interface{}) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	util.LogRequest(sp, in)

	var inputOrder *model.Order

	err := mapstructure.Decode(in, &inputOrder)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	// find cake first
	dataCake, err := e.cakeRepository.Get(sp, inputOrder.CakeId)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}
	util.LogObject(sp, "cake", dataCake)

	data, err := e.repository.InsertOrder(sp, inputOrder)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, data)
	return data, nil
}

func (e *OrderUsecase) GetAllOrder(ctx context.Context) (interface{}, error) {
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

func (e *OrderUsecase) GetOrder(ctx context.Context, id int) (interface{}, error) {
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

func (e *OrderUsecase) DeleteOrder(ctx context.Context, id int) error {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	_, err := e.repository.Get(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return errors.New("Order Not Fund")
	}

	err = e.repository.Delete(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return err
	}

	return nil
}

func (e *OrderUsecase) UpdateOrder(ctx context.Context, id int, in interface{}) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	if in == nil {
		return nil, errors.New("request cannot be nil")
	}

	var inputOrder *model.Order
	err := mapstructure.Decode(in, &inputOrder)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	_, err = e.repository.Get(sp, id)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("Order Not Fund")
	}

	// find cake first
	dataCake, err := e.cakeRepository.Get(sp, inputOrder.CakeId)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}
	util.LogObject(sp, "cake", dataCake)

	data, err := e.repository.Update(sp, id, inputOrder)
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}
	util.LogResponse(sp, data)

	return data, nil
}

func (e *OrderUsecase) CreateOrderBulk(ctx context.Context, in interface{}) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	util.LogRequest(sp, in)

	var inputOrder *entity.OrderInputBulk
	err := mapstructure.Decode(in, &inputOrder)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	var result []interface{}
	var resultError []string
	var wg sync.WaitGroup

	for _, v := range inputOrder.Data {
		wg.Add(1)
		go func(v entity.OrderInput) {
			defer wg.Done()

			// find cake first
			dataCake, err := e.cakeRepository.Get(sp, v.CakeId)
			if err != nil {
				util.LogError(sp, err)
				erro := errors.New(fmt.Sprintf("Cake not found in cake_id : %d", v.CakeId))
				resultError = append(resultError, erro.Error())
				return
			}
			util.LogObject(sp, "cake", dataCake)

			data, err := e.repository.InsertOrder(sp, &model.Order{
				CakeId: v.CakeId,
				Qty:    v.Qty,
			})
			if err != nil {
				util.LogError(sp, err)
				resultError = append(resultError, err.Error())
				return
			}
			result = append(result, data)
			util.LogObject(sp, "data order", data)
		}(v)
	}

	wg.Wait()

	response := make(map[string]interface{})
	response["success"] = result
	response["error"] = resultError

	return response, err

}
