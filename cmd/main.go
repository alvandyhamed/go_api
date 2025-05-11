package main

import (
	"fmt"
	"go_api_learn/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	config.LoadConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBSSLMode,
	)
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("  %v❌ اتصال به دیتابیس برقرار نشد:", err)

	}
	fmt.Println("✅ اتصال به دیتابیس برقرار شد.")

}
