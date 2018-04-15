package chats

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ChatRegister router register
func ChatRegister(router *gin.RouterGroup) {
	router.GET("/", GetAllChat)
	router.POST("/", SaveChat)
	router.GET("/:id", GetChatByID)
	router.PUT("/:id", UpdateChat)
	router.DELETE("/:id", DropChat)
}

// GetAllChat route handler chat
func GetAllChat(c *gin.Context) {
	chatList := FindAll()

	if len(chatList) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "chat is not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   chatList,
		})
	}
}

// SaveChat , route handler save chat data
func SaveChat(c *gin.Context) {
	var chat Chats
	if err := c.BindJSON(&chat); err == nil {
		returnMessage := chat.Save()

		c.JSON(http.StatusOK, gin.H{
			"status":  returnMessage.Status,
			"message": returnMessage.Description,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

// GetChatByID , route handler save chat data
func GetChatByID(c *gin.Context) {
	var chat, condition Chats
	chatID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = uint(chatID)
	chat.FindOne(condition)

	if chat.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "chat is not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   chat,
		})
	}
}

// UpdateChat , route handler update chat data
func UpdateChat(c *gin.Context) {
	var chat, data, condition Chats

	chatID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = uint(chatID)
	chat.FindOne(condition)

	if err := c.BindJSON(&data); err == nil {
		returnMessage := chat.Update(data)

		c.JSON(http.StatusOK, gin.H{
			"status":  returnMessage.Status,
			"message": returnMessage.Description,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

// DropChat , route handler update chat data
func DropChat(c *gin.Context) {
	var chat, condition Chats

	chatID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = uint(chatID)
	chat.FindOne(condition)

	returnMessage := chat.Drop()

	c.JSON(http.StatusOK, gin.H{
		"status":  returnMessage.Status,
		"message": returnMessage.Description,
	})
}
