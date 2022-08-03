package main

import (
	"fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)
type Data1 struct {
    gorm.Model
	title    string `json:"title"`
	content  string `json:"content"`
}

func main() {
    fmt.Println("Start migrate!");
    up01()
    down01();
    fmt.Println("End migrate!");
}

func up01() {
    fmt.Println("Start 001_up!");
    // dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		panic("failed to connect database")
	}
    // テーブル作成
    db.AutoMigrate(&Data1{})
    fmt.Println("End 001_up!");
}
func down01() {
    fmt.Println("Start 001_down!");
    // dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		panic("failed to connect database")
	}
    // テーブル削除
    db.Migrator().DropTable(&Data1{})
    fmt.Println("End 001_down!");
}