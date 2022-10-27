package main

import (
	"github.com/Wallaya007/sa-65-example/controller"
	"github.com/Wallaya007/sa-65-example/entity"
	"github.com/Wallaya007/sa-65-example/middlewares"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUser)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			// Title Routes
			protected.GET("/titles", controller.ListTitles)
			

			// Gender Routes
			protected.GET("/genders", controller.ListGenders)
			// Blood  Routes
			protected.GET("/bloods", controller.ListBloods)
			// Disease  Routes
			protected.GET("/diseases", controller.ListDiseases)
			// Patient  Routes
			protected.GET("/patients", controller.ListPatient)
			protected.POST("/patients", controller.CreatePatient)
			
		}
	}

	// Signup User Route
	r.POST("/createusers", controller.CreateUser)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
