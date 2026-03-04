package config

import (
	"todo-list/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:pswxxxxxxxx@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库结构
	err = db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
