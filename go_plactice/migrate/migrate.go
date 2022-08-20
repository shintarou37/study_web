package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

// 構造体名が小文字だと、初期化時にフィールドは存在しているがDB作成時にカラムが生成されていないので大文字にする必要がある。
type Data1 struct {
	gorm.Model
	Title   string
	Content string
}

func main() {
	fmt.Println("Start migrate!")
	dsn := "root:@tcp(127.0.0.1:3306)/go_plactice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(reflect.TypeOf(db))
	if err != nil {
		panic("failed to connect database")
	} else {
		down01(dsn, db)
		up01(dsn, db)
	}
	fmt.Println("End migrate!")
}

func up01(dsn string, db *gorm.DB) {
	fmt.Println("Start up01!")
	// charsetをutf8mb4にしないと、ORMをDBに接続した際のcharsetと合わずに文字列を登録すると「?」になる
	db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(Data1{})
	fmt.Println("End up01!")
}

func down01(dsn string, db *gorm.DB) {
	fmt.Println("Start down01!")
	// テーブル削除
	db.Migrator().DropTable(&Data1{})
	fmt.Println("End down01!")
}
