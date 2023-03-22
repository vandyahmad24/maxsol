package model

import (
	"gorm.io/gorm"
	"time"
)

type Cake struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName overrides the table name used by User to `profiles`
func (Cake) TableName() string {
	return "cake"
}

func (u *Cake) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()

	u.UpdatedAt = time.Now()
	return nil
}
