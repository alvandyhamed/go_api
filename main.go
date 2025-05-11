package main

import (
	"fmt"
	"go_api_learn/config"
	"go_api_learn/controller/swagger_handlers"
	_ "go_api_learn/docs"
	"go_api_learn/model"
	"go_api_learn/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// @title API پروژه ثبت‌نام
// @version 1.0
// @description مستندات API با Swagger

// @host localhost:8080
// @BasePath /

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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("  %v❌ اتصال به دیتابیس برقرار نشد:", err)

	}
	fmt.Println("✅ اتصال به دیتابیس برقرار شد.")

	authSwagger := &swagger_handlers.AuthSwaggerHandler{DB: db}
	userSwagger := &swagger_handlers.UserSwaggerHandler{DB: db}

	model.Migration(db)

	r := routes.SetupRouter(authSwagger, userSwagger)

	port := config.AppConfig.AppPort
	log.Printf("🚀 سرور در حال اجرا روی پورت %s", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("❌ خطا در اجرای سرور:", err)
	}

}
