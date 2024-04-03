package model

import "gorm.io/gorm"


type Project struct {
	Id    int    `gorm:"primaryKey"`
	Title string `gorm:"size:255;not null;unique" form:"title" binding:"required"`
	Image Image  `gorm:"foreignKey:ImageId"`
    ImageId int `gorm:"not null"`
    Description string `gorm:"size:255;not null" form:"description" binding:"required"`
    Url string `gorm:"size:255;not null" form:"url" binding:"required"`
    gorm.Model
}
