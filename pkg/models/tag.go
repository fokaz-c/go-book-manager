package models

import (
	"fmt"

	"github/fokaz-c/go-book-manager/pkg/config"

	"gorm.io/gorm"
)

type Tag struct {
	Tag   string `gorm:"unique;not null" json:"tag"`
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

func GetAllTags() []Tag {
	var books []Tag
	db.Find(&books)
	return books
}

func GetTagByID(tag string) (*Tag, *gorm.DB) {
	var getTag Tag
	db := db.Where("tag = ?", tag).Find(&getTag)
	return &getTag, db
}

func DeleteTag(tag string) Tag {
	var _tag Tag
	db.Where("tag = ?", tag).Delete(&_tag)
	return _tag
}
