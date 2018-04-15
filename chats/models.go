package chats

import (
	"../common"
	"github.com/jinzhu/gorm"
)

// Chats is chats model
type Chats struct {
	gorm.Model
	Message string `json:"Message"`
	UserID  uint   `json:"UserID"`
}

// AutoMigrate if not exist table
func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Chats{})
}

// Save chat data
func (chat *Chats) Save() common.ReturnMessage {
	db := common.GetDB()
	err := db.Save(&chat).Error

	return GenerateMessage(err)
}

// Update chat data
func (chat *Chats) Update(data interface{}) common.ReturnMessage {
	db := common.GetDB()
	err := db.Model(&chat).Update(data).Error

	return GenerateMessage(err)
}

// Drop chat data
func (chat *Chats) Drop() common.ReturnMessage {
	db := common.GetDB()
	err := db.Delete(&chat).Error

	return GenerateMessage(err)
}

// FindOne , select specific user
func (chat *Chats) FindOne(condition interface{}) {
	db := common.GetDB()
	db.Where(condition).First(&chat)
}

// FindAll , select all chat
func FindAll() []Chats {
	var chats []Chats

	db := common.GetDB()
	db.Find(&chats)

	return chats
}

// FindMany select by where dynamic data
func FindMany(condition interface{}) []Chats {
	var chats []Chats

	db := common.GetDB()
	db.Where(condition).Find(&chats)

	return chats
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
