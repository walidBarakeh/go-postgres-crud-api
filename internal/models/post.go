package models

import (
    "time"
)

type Post struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title" gorm:"not null"`
    Content   string    `json:"content"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Published bool      `json:"published" gorm:"default:false"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
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
