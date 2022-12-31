package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"model"`
	Role       int    `json:"role"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}
