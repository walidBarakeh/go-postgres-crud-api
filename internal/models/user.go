package models

import (
    "gorm.io/gorm"
)

type User struct {
	gorm.Model
    Name      string    `json:"name" gorm:"not null"`
    Email     string    `json:"email" gorm:"uniqueIndex;not null"`
    Age       *int      `json:"age,omitempty"`
    Posts     []Post    `json:"posts,omitempty" gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   *int   `json:"age,omitempty"`
}

type UpdateUserRequest struct {
    Name  string `json:"name,omitempty"`
    Email string `json:"email,omitempty"`
    Age   *int   `json:"age,omitempty"`
}
