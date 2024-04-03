package model

import (
	"apallis/portfolio/database"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"primaryKey"`
	Username string `gorm:"size:255;not null;unique" form:"username" binding:"required"`
	Email    string `gorm:"size:255;not null;unique" form:"email"`
	Password string `gorm:"size:255;not null" form:"password" binding:"required"`
	gorm.Model
}

func GetUserByEmailAndPassword(email, password string) (User, error) {
	// TODO: This need to compare the password
	var user User
	result := database.DB.Where("email = ? AND password = ?", email, password).First(&user)
	return user, result.Error
}

func GetUserByUsernameAndPassword(username, password string) (User, error) {
	var user User
	result := database.DB.Where("username = ?", username).First(&user)
	if !comparePassword(user.Password, password) {
		user = User{}
		return user, fmt.Errorf("Invalid password")
	}
	return user, result.Error
}

func comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
