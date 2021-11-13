package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This middleware ensures that a request will be aborted with an error
// If the user is not logged in
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			// if token, err := c.Cookie("token"); err == nil || token != "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
