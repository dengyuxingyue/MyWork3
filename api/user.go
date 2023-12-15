package api

import (
	"net/http"
	"work/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
