// models.go
package main

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Author    string         `gorm:"not null" json:"author"`
	Content   string         `gorm:"not null" json:"content"`
	Comments  []Comment      `json:"comments"` // foreignKey:ArticleID inferred by GORM
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Comment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	ArticleID uint           `gorm:"not null" json:"article_id"`
	Name      string         `gorm:"not null" json:"name"`
	Comment   string         `gorm:"not null" json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
