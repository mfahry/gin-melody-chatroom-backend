package main

import (
	"./chatroom"
	"./chats"
	"./common"
	"./users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	// m := melody.New()

	// init schema
	common.Init()

	// migrate or create table of struct
	users.AutoMigrate()
	chats.AutoMigrate()

	// add router
	users.UserRegister(router.Group("/user"))
	chats.ChatRegister(router.Group("/chat"))
	chatroom.ChatroomRegister(router.Group("/chatroom"))

	router.Run()
}
