package entity

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model

	ID        int    `json:"-" gorm:"primaryKey;auto_increment"`
	Title     string `json:"title" binding:"required" gorm:"type:varchar(65);UNIQUE"`
	Semester  int    `json:"semester" binding:"lte=1;gte=8"`
	Credits   int    `json:"credits" binding:"lte=1;gte=6"`
	Mandatory bool   `json:"mandatory"`
}
