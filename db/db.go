package db

import (
	"github.com/Aj002Th/LittlePrince/data"
	"github.com/Aj002Th/LittlePrince/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = data.GetMySQLConn()

	// 迁移 & 建表
	db.AutoMigrate(&model.User{})
}
