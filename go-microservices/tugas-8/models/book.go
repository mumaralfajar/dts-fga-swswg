package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	// *gorm.Model
	ID        int       `gorm:"primaryKey" json:"id"`
	Name_Book string    `gorm:"not null;type:varchar(100)" json:"name_book" binding:"required"`
	Author    string    `gorm:"not null;type:varchar(100)" json:"author" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// method || hooks
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if len(b.Name_Book) < 2 {
		err = errors.New("title is too short")
	}
	if len(b.Author) < 2 {
		err = errors.New("author is too short")
	}

	return
}