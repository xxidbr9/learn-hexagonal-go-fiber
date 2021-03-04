package bookquery

import (
	"errors"

	"gorm.io/gorm"
)

// GetBooks :
func GetBooks(db *gorm.DB) []BookQuery {
	var Books []BookQuery

	db.Find(&Books)
	if len(Books) == 0 {
		return []BookQuery{}
	}
	return Books
}

// GetBookByID :
func GetBookByID(db *gorm.DB, param string) (BookQuery, error) {
	var book BookQuery

	result := db.Find(&book, param)
	if result.RowsAffected == 0 {
		return book, errors.New("Book not found")
	}
	return book, nil
}
