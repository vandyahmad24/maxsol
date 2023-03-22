package cake_repository

import (
	"errors"
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/model"
	"vandyahmad24/maxsol/app/util"

	"github.com/opentracing/opentracing-go"
)

type Cake struct {
	db *gorm.DB
}

func NewCake(db *gorm.DB) *Cake {
	return &Cake{
		db: db,
	}
}

func (cl *Cake) InsertCake(span opentracing.Span, input *model.Cake) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "InsertCake")
	defer sp.Finish()
	util.LogRequest(sp, input)

	err := cl.db.Table("cake").Create(input).Error
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, input)
	return input, nil
}

func (cl *Cake) GetAll(span opentracing.Span) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "GetAll")
	defer sp.Finish()

	var cakes []model.Cake
	err := cl.db.Table("cake").Order("rating DESC, title").Find(&cakes).Error
	if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, cakes)
	return cakes, nil
}

func (cl *Cake) Get(span opentracing.Span, id int) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "Get")
	defer sp.Finish()

	var cake model.Cake
	err := cl.db.Table("cake").First(&cake, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Cake Not Found")
	} else if err != nil {
		util.LogError(sp, err)
		return nil, err
	}

	util.LogResponse(sp, cake)
	return cake, nil
}

func (cl *Cake) Delete(span opentracing.Span, id int) error {
	sp := util.CreateSubChildSpan(span, "Delete")
	defer sp.Finish()

	err := cl.db.Table("cake").Delete(&model.Cake{}, id).Error
	if err == gorm.ErrRecordNotFound {
		util.LogError(sp, err)
		return errors.New("Failed Delete")
	}

	return nil
}

func (cl *Cake) Update(span opentracing.Span, id int, input *model.Cake) (interface{}, error) {
	sp := util.CreateSubChildSpan(span, "Update")
	defer sp.Finish()

	var cake model.Cake

	err := cl.db.Table("cake").Where("id = ?", id).Updates(&model.Cake{
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Price:       input.Price,
	}).Scan(&cake).Error
	if err != nil {
		util.LogError(sp, err)
		return nil, errors.New("Failed Update")
	}

	util.LogResponse(sp, cake)
	return cake, nil
}
