package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	todoapp "github.com/dkhoanguyen/todo-app"
	"github.com/dkhoanguyen/todo-app/models"
	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	todoapp.Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func ShowRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	todoapp.Render(c, gin.H{
		"title": "Register"}, "register.html")
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func PerformLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	if models.IsUserValid(username, password) {
		token := generateSessionToken()
		c.SetSameSite(sameSiteCookie)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		todoapp.Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func PerformLogout(c *gin.Context) {
	var sameSiteCookie http.SameSite
	c.SetSameSite(sameSiteCookie)
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		token := generateSessionToken()
		c.SetSameSite(sameSiteCookie)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		todoapp.Render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}
