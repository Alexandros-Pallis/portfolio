package model

import (
	"apallis/portfolio/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PermissionType string

const (
	Read   PermissionType = "read"
	Write  PermissionType = "write"
	Delete PermissionType = "delete"
)

type Permission struct {
	Id   int            `gorm:"primaryKey"`
	Name PermissionType `gorm:"size:255;not null;unique" form:"name" binding:"required"`
	gorm.Model
}

func (permission *Permission) GetById(id string) error {
	result := database.DB.First(&permission, id)
	return result.Error
}

func (permission *Permission) GetAll() ([]Permission, error) {
	var permissions []Permission
	result := database.DB.Find(&permissions)
	return permissions, result.Error
}

type User struct {
	Id          int          `gorm:"primaryKey"`
	Username    string       `gorm:"size:255;not null;unique" form:"username" binding:"required"`
	Email       string       `gorm:"size:255;not null;unique" form:"email"`
	Password    string       `gorm:"size:255;not null" form:"password" binding:"required"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
	gorm.Model
}

func (user *User) HasPermission(permissionType PermissionType) bool {
	for _, permission := range user.Permissions {
		if permission.Name == permissionType {
			return true
		}
	}
	return false
}

func (user *User) AddPermission(permission Permission) {
	user.Permissions = append(user.Permissions, permission)
}

func (user *User) RemovePermissions() error {
	return database.DB.Model(&user).Association("Permissions").Clear()
}

func (user *User) GetPermissionsAsString() string {
	permissions := ""
	for _, permission := range user.Permissions {
		permissions += string(permission.Name) + ", "
	}
	if len(permissions) == 0 {
		permissions = "No permissions added yet..."
	} else {
        permissions = permissions[:len(permissions)-2]
    }
	return permissions
}

func (user *User) GetAll() ([]User, error) {
	var users []User
	err := database.DB.Model(&User{}).Preload("Permissions").Find(&users).Error
	return users, err
}

func (user *User) GetCurrentUser(contenxt *gin.Context) User {
    var currentUser User
    contextUser, ok := contenxt.Get("user")
    if !ok {
        return currentUser
    }
    currentUser = contextUser.(User)
    return currentUser
}

func GetUserByUsernameAndPassword(username, password string) (User, error) {
	var user User
    err := database.DB.Preload("Permissions").Where("username = ?", username).First(&user).Error
	if !comparePassword(user.Password, password) {
		user = User{}
		return user, fmt.Errorf("Invalid password")
	}
	return user, err
}

func comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (user *User) Save() error {
	fmt.Println("user perms: ", user.GetPermissionsAsString())
	result := database.DB.Save(user)
	return result.Error
}
