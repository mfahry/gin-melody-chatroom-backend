package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRegister router register
func UserRegister(router *gin.RouterGroup) {
	router.GET("/", GetAllUser)
	router.POST("/", SaveUser)
	router.GET("/:id", GetUserByID)
	router.PUT("/:id", UpdateUser)
	router.DELETE("/:id", DropUser)
}

// GetAllUser route handler user
func GetAllUser(c *gin.Context) {
	userList := FindAll()

	if len(userList) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "user is not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   userList,
		})
	}
}

// SaveUser , route handler save user data
func SaveUser(c *gin.Context) {
	var user Users
	if err := c.BindJSON(&user); err == nil {
		returnMessage := user.Save()

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

// GetUserByID , route handler save user data
func GetUserByID(c *gin.Context) {
	var user, condition Users

	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = uint(userID)
	user.FindOne(condition)

	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "user is not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   user,
		})
	}
}

// UpdateUser , route handler update user data
func UpdateUser(c *gin.Context) {
	var user, data, condition Users

	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = uint(userID)
	user.FindOne(condition)

	if err := c.BindJSON(&data); err == nil {
		returnMessage := user.Update(data)

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

// DropUser , route handler update user data
func DropUser(c *gin.Context) {
	var user, condition Users

	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = uint(userID)
	user.FindOne(condition)

	returnMessage := user.Drop()

	c.JSON(http.StatusOK, gin.H{
		"status":  returnMessage.Status,
		"message": returnMessage.Description,
	})
}
