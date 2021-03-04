package book

import "gorm.io/gorm"

// NewBookApp :
func NewBookApp(db *gorm.DB) error {
	return db.AutoMigrate(&BookModel{})
}
