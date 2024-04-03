package model

import "gorm.io/gorm"

type Skill struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null;unique" form:"name" binding:"required"`
	Icon        Image  `gorm:"foreignKey:IconId"`
    IconId      int    `gorm:"not null"`
	gorm.Model
}
