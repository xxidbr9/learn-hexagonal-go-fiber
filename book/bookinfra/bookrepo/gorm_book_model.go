package bookrepo

import (
	"nando/gorm_fiber/shared"
)

// BookRepo :
type BookRepo struct {
	shared.Model
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint    `json:"quantity"`
}

// TableName : Rename table name
func (book *BookRepo) TableName() string {
	return "tb_buku"
}
