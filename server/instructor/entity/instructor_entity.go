package entity

import (
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model

	ID          int    `json:"-" gorm:"primaryKey;auto_increment"`
	FirstName   string `json:"first_name" binding:"required" gorm:"type:varchar(20)"`
	LastName    string `json:"last_name" binding:"required" gorm:"type:varchar(20)"`
	PhoneNumber string `json:"phone_number" binding:"gte=9,lte9" gorm:"type:varchar(9);UNIQUE"`
	Email       string `json:"email" binding:"required" gorm:"type:varchar(60);UNIQUE"`
	LinkedIn    string `json:"linked_in" binding:"url" gorm:"type:varchar(256);UNIQUE"`
	Rating      int    `json:"rating"`
}
