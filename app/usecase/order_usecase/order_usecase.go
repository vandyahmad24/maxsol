package order_usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
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

			util.LogObject(sp, "data order", data)
		}(v)
	}

	wg.Wait()

	response := make(map[string]interface{})
	response["error_data"] = resultError
	response["total_data"] = len(inputOrder.Data)
	response["total_success"] = len(inputOrder.Data) - len(resultError)

	return response, err

}

func (e *OrderUsecase) CreateOrderBulkWithWorker(ctx context.Context, in interface{}) (interface{}, error) {
	sp := util.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	util.LogRequest(sp, in)

	var inputOrder *entity.OrderInputBulk
	err := mapstructure.Decode(in, &inputOrder)
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	jobs := make(chan entity.OrderInput, 0)
	wg := new(sync.WaitGroup)
	errChan := make(chan error, len(inputOrder.Data))

	go e.dispatchWorkers(ctx, jobs, wg, errChan)
	e.readData(inputOrder, jobs, wg)
	var resultError []string
	go func() {
		for err := range errChan {
			if err != nil {
				wg.Add(1)
				fmt.Println("error chan ", err)
				resultError = append(resultError, err.Error())
				wg.Done()
			}
		}
	}()

	wg.Wait()
	close(errChan)

	response := make(map[string]interface{})
	response["error_data"] = resultError
	response["total_data"] = len(inputOrder.Data)
	response["total_success"] = len(inputOrder.Data) - len(resultError)

	return response, err

}

func (e *OrderUsecase) dispatchWorkers(ctx context.Context, jobs <-chan entity.OrderInput, wg *sync.WaitGroup, errChan chan<- error) {
	for workerIndex := 0; workerIndex <= 10; workerIndex++ {
		go func(ctx context.Context, workerIndex int, jobs <-chan entity.OrderInput, wg *sync.WaitGroup, errChan chan<- error) {
			counter := 0

			for job := range jobs {
				if err := e.doTheJob(ctx, workerIndex, counter, job); err != nil {
					errChan <- err
				}
				wg.Done()
				counter++
			}
		}(ctx, workerIndex, jobs, wg, errChan)
	}
}

func (e *OrderUsecase) doTheJob(ctx context.Context, workerIndex, counter int, v entity.OrderInput) error {
	sp := util.CreateChildSpan(ctx, string("doTheJob"))
	defer sp.Finish()
	dataCake, err := e.cakeRepository.Get(sp, v.CakeId)
	if err != nil {
		util.LogError(sp, err)
		err = errors.New(fmt.Sprintf("Cake not found in cake_id : %d", v.CakeId))
		//resultError = append(resultError, erro.Error())
		return err
	}
	util.LogObject(sp, "cake", dataCake)

	data, err := e.repository.InsertOrder(sp, &model.Order{
		CakeId: v.CakeId,
		Qty:    v.Qty,
	})
	if err != nil {
		util.LogError(sp, err)
		//resultError = append(resultError, err.Error())
		return err
	}
	//result = append(result, data)
	util.LogObject(sp, "data order", data)

	if counter%100 == 0 {
		log.Println("=> worker", workerIndex, "inserted", counter, "data")
	}
	return nil
}

func (e *OrderUsecase) readData(input *entity.OrderInputBulk, jobs chan<- entity.OrderInput, wg *sync.WaitGroup) {

	for _, v := range input.Data {
		wg.Add(1)
		jobs <- v
	}

	close(jobs)
}
