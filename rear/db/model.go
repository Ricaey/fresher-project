package db


// 评论模型
type Comment struct {
    ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name    string `gorm:"type:varchar(64);not null" json:"name"`
    Content string `gorm:"type:text;not null" json:"content"`
}

 