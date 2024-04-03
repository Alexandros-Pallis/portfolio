package model

import "gorm.io/gorm"

type Image struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null;unique" form:"name" binding:"required"`
	Alt  string `gorm:"size:255;not null" form:"alt" binding:"required"`
	Path string `gorm:"size:255;not null" form:"path" binding:"required"`
	gorm.Model
}
