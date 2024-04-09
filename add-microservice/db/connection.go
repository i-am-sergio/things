package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// mysql://root:wfLrySrJTJmcpoDgCryIShnfcgNrJMEW@viaduct.proxy.rlwy.net:47496/railway
	DSN := "root:wfLrySrJTJmcpoDgCryIShnfcgNrJMEW@tcp(viaduct.proxy.rlwy.net:47496)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var error error
	DB, error = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if error != nil {
		panic("Failed to connect to database!")
	} else {
		println("Connected to database: ", DB.Name())
	}
}
