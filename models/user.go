package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password string `json:"password" gorm:"size:100;not null"`
	Email    string `json:"email" gorm:"size:100"`
	Status   int    `json:"status" gorm:"default:1"` // 1: 正常, 0: 禁用
	Todos    []Todo `json:"todos" gorm:"foreignKey:UserID"`
}