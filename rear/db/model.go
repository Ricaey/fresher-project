package db

import "gorm.io/gorm"

// 评论模型
type Comment struct {
    ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name    string `gorm:"type:varchar(64);not null" json:"name"`
    Content string `gorm:"type:text;not null" json:"content"`
}

// 用户模型（如需扩展用户功能，可用）
type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex;type:varchar(64);not null"`
}