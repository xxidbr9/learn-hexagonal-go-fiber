package book

import "nando/gorm_fiber/shared"

// BookModel :
type BookModel struct {
	shared.Model
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint    `json:"quantity"`
}
func (book *BookModel) TableName() string {
	return "tb_buku"
}
