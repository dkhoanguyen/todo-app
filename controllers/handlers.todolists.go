package controllers

import (
	"net/http"
	"strconv"

	todoapp "github.com/dkhoanguyen/todo-app"
	"github.com/dkhoanguyen/todo-app/models"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	todos := models.GetAllTodos()

	todoapp.Render(c, gin.H{
		"title":   "Home Page",
		"payload": todos}, "index.html")
}

func ShowTodoCreationPage(c *gin.Context) {
	todoapp.Render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func GetTodo(c *gin.Context) {
	if todoID, err := strconv.Atoi(c.Param("todo_id")); err == nil {
		if todo, err := models.GetTodoByID(todoID); err == nil {
			todoapp.Render(c, gin.H{"title": todo.Title,
				"payload": todo}, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if t, err := models.CreateNewTodo(title, content); err == nil {
		todoapp.Render(c, gin.H{
			"title":   "Submission Successful",
			"payload": t}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
