package routes

import (
	"crud-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)

			userByID := users.Group("/:id")
			{
				userByID.GET("", handlers.GetUser)
				userByID.PUT("", handlers.UpdateUser)
				userByID.DELETE("", handlers.DeleteUser)
				

				userByID.GET("/posts", handlers.GetUserPosts)
			}
		}

		posts := api.Group("/posts")
		{
			posts.GET("", handlers.GetPosts)
			posts.POST("", handlers.CreatePost)
			posts.GET("/:id", handlers.GetPost)
			posts.PUT("/:id", handlers.UpdatePost)
			posts.DELETE("/:id", handlers.DeletePost)
		}
	}
	return r
}