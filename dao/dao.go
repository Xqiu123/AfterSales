package dao

import (
	"as/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init dao
func Init() {
	// init db
	model.DB.Init()
	DB = model.DB.Self
}
