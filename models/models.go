package models

import (
	"github.com/lib/pq"
	"github.com/syed-ashraf123/go-crud/database"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string         `json:"title" validate:"required"`
	Body    string         `json:"body" validate:"required"`
	TagList pq.StringArray `json:"tagList" gorm:"type:text[]"`
}

func Migrate() {
	database.DB.AutoMigrate(&Article{})
}
