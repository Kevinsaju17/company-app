package routes

import (
	"company-app/backend/controllers"
	"company-app/backend/middleware"

	"github.com/gin-gonic/gin"
)


// Register protected routes
func ProfileRoutes(r *gin.Engine) {

	// Protected GET route
	r.GET(
		"/profile",

		// Middleware executes BEFORE route
		middleware.AuthMiddleware(),

		// Actual controller
		controllers.Profile,
	)
}