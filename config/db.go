package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Yigit4434.@tcp(127.0.0.1:3306)/taskmanager_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Veritabanına bağlanılamadı: ", err)
	}

	fmt.Println("✅ Veritabanı bağlantısı başarılı.")
	DB = database
}
