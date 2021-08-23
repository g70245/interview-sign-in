package router

import (
	"app/controller"
	"app/model"
	"app/service"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	// Inject depedencies
	var userStore model.UserStore = model.UserStoreMock{}
	userController := controller.UserController{UserManager: &service.UserManager{UserStore: &userStore}}

	r.POST("/signin", userController.Signin)

	return r
}
