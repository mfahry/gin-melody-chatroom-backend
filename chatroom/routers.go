package chatroom

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../chats"
	"../users"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var m = melody.New()

// ChatroomRegister is register router
func ChatroomRegister(router *gin.RouterGroup) {
	router.POST("/login", LoginChatRoom)
	router.GET("/history-chat", HistoryChat)
	router.GET("/ws", HandleRequest)

	m.HandleMessage(func(s *melody.Session, message []byte) {
		var chat chats.Chats
		err := json.Unmarshal(message, chat)

		if err != nil {
			fmt.Println("error : ", err)
		} else {
			chat.Save()
			m.Broadcast([]byte(chat.Message))
		}
	})
}

// LoginChatRoom is function to login
func LoginChatRoom(c *gin.Context) {
	var user, condition users.Users

	if err := c.BindJSON(&condition); err == nil {
		user.FindOne(condition)

		if user.ID == 0 {
			returnMessage := condition.Save()
			if returnMessage.Status != 200 {
				c.JSON(returnMessage.Status, gin.H{
					"status":  returnMessage.Status,
					"message": returnMessage.Description,
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "login success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

// HistoryChat -> get all chat by user
func HistoryChat(c *gin.Context) {
	chatList := chats.FindAll()

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"chats":  chatList,
	})
}

// HandleRequest Web Socket
func HandleRequest(c *gin.Context) {
	m.HandleRequest(c.Writer, c.Request)
}
