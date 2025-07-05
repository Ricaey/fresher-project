package db

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 读取 config.json 并返回 DSN
func getDSNFromConfig() (string, error) {
	type myConfig struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	}
	type config struct {
		Database myConfig `json:"database"`
	}
	f, err := os.Open("../config.json")
	if err != nil {
		return "", err
	}
	defer f.Close()
	var cfg config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return "", err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	return dsn, nil
}

// 初始化数据库
func InitDB() error {
	dsn, err := getDSNFromConfig()
	if err != nil {
		return fmt.Errorf("读取配置失败: %w", err)
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 这里改为mysql.Open
	if err != nil {
		return fmt.Errorf("数据库连接失败: %w", err)
	}
	// 自动迁移
	if err := DB.AutoMigrate(&Comment{}); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}
	fmt.Println("数据库连接成功")
	return nil
}