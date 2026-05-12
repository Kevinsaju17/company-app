package controllers

import (
	"company-app/backend/config"
	"company-app/backend/models"
	"company-app/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


// Validator instance
var validate = validator.New()


// REGISTER USER
func Register(c *gin.Context) {

	var user models.User

	// Read JSON request body
	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Validate request fields
	if err := validate.Struct(user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Hash plain password before saving
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})

		return
	}

	// Replace plain password with hashed password
	user.Password = hashedPassword

	// Save user into PostgreSQL
	result := config.DB.Create(&user)

	// DB insertion failed
	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User creation failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}



// LOGIN USER
func Login(c *gin.Context) {

	var input models.User

	// Read login request JSON
	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var user models.User

	// Search user by email
	result := config.DB.Where("email = ?", input.Email).First(&user)

	// User not found
	if result.Error != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare entered password with hashed DB password
	valid := utils.CheckPassword(input.Password, user.Password)

	if !valid {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.Email)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token generation failed",
		})

		return
	}

	// Return token to client
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}