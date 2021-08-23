package controller

import (
	"app/helper/status"
	"app/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type UserController struct {
	UserManager *service.UserManager
}

type SigninRequest struct {
	Username string `json:"username" validate:"min=4,max=16,regexp=^[a-zA-Z0-9]*$"`
	Password string `json:"password" validate:"min=8,max=16"`
}

func (controller *UserController) Signin(c *gin.Context) {
	var signinRequest SigninRequest
	if err := c.ShouldBindJSON(&signinRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": status.InternalError,
		})
		return
	}

	if err := validator.Validate(signinRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": status.InputFormatError,
		})
		return
	}

	if err := controller.UserManager.Signin(signinRequest.Username, signinRequest.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": status.Success,
	})
}
