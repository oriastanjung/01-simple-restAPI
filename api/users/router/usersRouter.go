package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oriastanjung/01-simple-restAPI/api/users/controller"
)

func UsersGroup(router *gin.Engine) {

	usersGroup := router.Group("/users")

	{
		usersGroup.GET("/", controller.GetAllUsers())
		usersGroup.POST("/", controller.CreateOneUser())
	}
}
