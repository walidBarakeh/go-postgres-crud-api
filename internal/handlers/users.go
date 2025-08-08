package handlers

import (
    "net/http"
    "strconv"

    "crud-api/internal/database"
    "crud-api/internal/models"

    "github.com/gin-gonic/gin"
)

// GetUsers retrieves all users
func GetUsers(c *gin.Context) {
    var users []models.User
    
    result := database.DB.Find(&users)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser retrieves a single user by ID
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    result := database.DB.Preload("Posts").First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
    var req models.CreateUserRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{
        Name:  req.Name,
        Email: req.Email,
        Age:   req.Age,
    }

    result := database.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": user})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    var req models.UpdateUserRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := database.DB.First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    database.DB.Model(&user).Updates(req)
    c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    result := database.DB.First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
