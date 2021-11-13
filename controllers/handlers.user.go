package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	if models.isUserValid(username, password) {

	}
}
