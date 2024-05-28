package models

import (
	"fmt"

	"github/fokaz-c/go-book-manager/pkg/config"
)

type Tag struct {
	Name  string `gorm:"unique;not null" json:"name"`
	Books string `json:"books"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&Tag{}); err != nil {
		fmt.Printf("Failed to migrate Tag schema: %v\n", err)
	}
}

func (t *Tag) CreateTag() (*Tag, error) {
	result := db.Create(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}
