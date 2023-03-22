package order_repository

import (
	"errors"
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/domain/entity"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"

	"github.com/opentracing/opentracing-go"
)

type Order struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) *Order {
	return &Order{
		db: db,
	}
}

func (o *Order) InsertOrder(span opentracing.Span, input *model.Order) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "InsertOrder")
	defer sp.Finish()
	util.LogRequest(sp, input)

	err := o.db.Table("orders").Create(input).Error
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	res := entity.OrderResponseWithoutCake{
		Id:        input.Id,
		CakeId:    input.CakeId,
		Qty:       input.Qty,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	util.LogResponse(sp, res)
	return res, nil
}

func (o *Order) GetAll(span opentracing.Span) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "GetAll")
	defer sp.Finish()

	var orders []model.Order
	err := o.db.Preload("Cake").Table("orders").Order("created_at DESC").Find(&orders).Error
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, orders)
	return orders, nil
}

func (o *Order) Get(span opentracing.Span, id int) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "Get")
	defer sp.Finish()

	var order model.Order
	err := o.db.Preload("Cake").Table("orders").Debug().First(&order, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Order Not Found")
	} else if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, order)
	return order, nil
}

func (o *Order) Delete(span opentracing.Span, id int) error {
	sp := util.CreateSubChildSpan(span, "Delete")
	defer sp.Finish()

	err := o.db.Table("orders").Delete(&model.Order{}, id).Error
	if err == gorm.ErrRecordNotFound {
		util.LogError(sp, err)
		return errors.New("Failed Delete")
	}

	return nil
}

func (o *Order) Update(span opentracing.Span, id int, input *model.Order) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "Update")
	defer sp.Finish()

	var order entity.OrderResponseWithoutCake

	err := o.db.Table("orders").Where("id = ?", id).Updates(&model.Order{
		CakeId: input.CakeId,
		Qty:    input.Qty,
	}).Scan(&order).Error
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("Failed Update")
	}

	util.LogResponse(sp, order)
	return order, nil
}
