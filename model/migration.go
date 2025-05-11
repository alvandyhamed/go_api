package model

import (
	"gorm.io/gorm"
	"log"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{})

	if err != nil {
		log.Fatal("❌ خطا در اجرای مایگریشن:", err)
	}
	log.Println("✅ مایگریشن با موفقیت انجام شد.")
}
