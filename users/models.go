package users

import (
	"../chats"
	"../common"
	"github.com/jinzhu/gorm"
)

// Users for user_model table
type Users struct {
	gorm.Model
	Username string        `json:"Username"`
	Fullname string        `json:"Fullname"`
	Chats    []chats.Chats `gorm:"foreignkey:UserID"`
}

// AutoMigrate if not exist table
func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Users{})
}

// Save user data
func (user *Users) Save() common.ReturnMessage {
	db := common.GetDB()
	err := db.Save(&user).Error

	return GenerateMessage(err)
}

// Update user data
func (user *Users) Update(data interface{}) common.ReturnMessage {
	db := common.GetDB()
	err := db.Model(&user).Update(data).Error

	return GenerateMessage(err)
}

// Drop user data
func (user *Users) Drop() common.ReturnMessage {
	db := common.GetDB()
	err := db.Delete(&user).Error

	return GenerateMessage(err)
}

// FindOne , select specific user
func (user *Users) FindOne(condition interface{}) {
	db := common.GetDB()
	db.Where(condition).First(&user)
}

// FindAll , select all user
func FindAll() []Users {
	var users []Users

	db := common.GetDB()
	db.Find(&users)

	return users
}

// FindMany select by where dynamic data
func FindMany(condition interface{}) []Users {
	var users []Users

	db := common.GetDB()
	db.Where(condition).Find(&users)

	return users
}

// GenerateMessage , generate message from execute
func GenerateMessage(err error) common.ReturnMessage {
	var returnMessage common.ReturnMessage

	if err != nil {
		returnMessage.Status = 500
		returnMessage.Description = err.Error()
	} else {
		returnMessage.Status = 200
		returnMessage.Description = `the process is success`
	}

	return returnMessage
}
