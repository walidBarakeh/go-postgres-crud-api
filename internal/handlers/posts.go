package handlers

import (
    "net/http"
    "strconv"

    "crud-api/internal/database"
    "crud-api/internal/models"

    "github.com/gin-gonic/gin"
)

// GetPosts retrieves all posts
func GetPosts(c *gin.Context) {
    var posts []models.Post
    
    result := database.DB.Preload("User").Find(&posts)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GetPost retrieves a single post by ID
func GetPost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    result := database.DB.Preload("User").First(&post, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": post})
}

// GetUserPosts retrieves all posts by a specific user
func GetUserPosts(c *gin.Context) {
    userID := c.Param("userId")
    var posts []models.Post

    result := database.DB.Where("user_id = ?", userID).Preload("User").Find(&posts)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": posts})
}

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
    var req models.CreatePostRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Verify user exists
    var user models.User
    if err := database.DB.First(&user, req.UserID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }

    post := models.Post{
        Title:   req.Title,
        Content: req.Content,
        UserID:  req.UserID,
    }

    if req.Published != nil {
        post.Published = *req.Published
    }

    result := database.DB.Create(&post)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    // Load the user relationship
    database.DB.Preload("User").First(&post, post.ID)

    c.JSON(http.StatusCreated, gin.H{"data": post})
}

// UpdatePost updates an existing post
func UpdatePost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post
    var req models.UpdatePostRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := database.DB.First(&post, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    database.DB.Model(&post).Updates(req)
    database.DB.Preload("User").First(&post, post.ID)
    
    c.JSON(http.StatusOK, gin.H{"data": post})
}

// DeletePost deletes a post
func DeletePost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    result := database.DB.First(&post, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    database.DB.Delete(&post)
    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

---

// internal/routes/routes.go
package routes

import (
    "crud-api/internal/handlers"
    
    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    api := r.Group("/api/v1")
    {
        // User routes
        users := api.Group("/users")
        {
            users.GET("", handlers.GetUsers)
            users.POST("", handlers.CreateUser)
            users.GET("/:id", handlers.GetUser)
            users.PUT("/:id", handlers.UpdateUser)
            users.DELETE("/:id", handlers.DeleteUser)
            users.GET("/:userId/posts", handlers.GetUserPosts)
        }

        // Post routes
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
