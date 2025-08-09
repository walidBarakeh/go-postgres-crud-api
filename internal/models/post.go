package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
    Title     string    `json:"title" gorm:"not null"`
    Content   string    `json:"content"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Published bool      `json:"published" gorm:"default:false"`
}

type CreatePostRequest struct {
    Title     string `json:"title" binding:"required"`
    Content   string `json:"content"`
    UserID    uint   `json:"user_id" binding:"required"`
    Published *bool  `json:"published,omitempty"`
}

type UpdatePostRequest struct {
    Title     string `json:"title,omitempty"`
    Content   string `json:"content,omitempty"`
    Published *bool  `json:"published,omitempty"`
}
