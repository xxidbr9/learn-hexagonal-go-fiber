package bookquery

import "nando/gorm_fiber/shared"

// BookQuery :
type BookQuery struct {
	shared.Model
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint    `json:"quantity"`
}


// TableName :
func (book *BookQuery) TableName() string {
	return "tb_buku"
}
