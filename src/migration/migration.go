package migration

import (
	"apallis/portfolio/model"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.Image{})
    db.AutoMigrate(&model.Project{})
    db.AutoMigrate(&model.Skill{})
}
