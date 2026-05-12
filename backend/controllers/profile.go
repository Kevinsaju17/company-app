package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// Protected profile route
func Profile(c *gin.Context) {

	// Get email stored by middleware
	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"message": "Protected route accessed",
		"user": email,
	})
}