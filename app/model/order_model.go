package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Id        int       `json:"id"`
	CakeId    int       `json:"cake_id"`
	Qty       int       `json:"qty"`
	Cake      *Cake     `gorm:"foreignKey:cake_id;references:id" json:"cake"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()

	u.UpdatedAt = time.Now()
	return nil
}
