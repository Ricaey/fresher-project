package db

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 读取 config.json 并返回 DSN
func getDSNFromConfig() (string, error) {
	type pgConfig struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	}
	type config struct {
		Database pgConfig `json:"database"`
	}
	f, err := os.Open("../config/config.json")
	if err != nil {
		return "", err
	}
	defer f.Close()
	var cfg config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return "", err
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
	return dsn, nil
}

// 初始化数据库
func InitDB() error {
	dsn, err := getDSNFromConfig()
	if err != nil {
		return fmt.Errorf("读取配置失败: %w", err)
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库连接失败: %w", err)
	}
	// 自动迁移
	if err := DB.AutoMigrate(&Comment{}, &User{}); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}
	fmt.Println("数据库连接成功")
	return nil
}