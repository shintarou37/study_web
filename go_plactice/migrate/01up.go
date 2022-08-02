package main

import (
	"fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func Up001() {
    fmt.Println("Start 001_up!");
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		panic("failed to connect database")
	}
    // テーブル作成
    db.AutoMigrate(&Data1{})
    fmt.Println("End 001_up!");
}