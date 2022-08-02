package main

import (
	"fmt"
	"net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)
type Data1 struct {
    gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func main() {
    fmt.Println("Start!");
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		panic("failed to connect database")
	}
    // テーブル作成
    db.AutoMigrate(&Data1{})
    // テーブル削除
    db.Migrator().DropTable(&Data1{})
    http.HandleFunc("/", handler);
    http.ListenAndServe(":8080", nil)
    fmt.Println("End!");

}
func handler(w http.ResponseWriter, r *http.Request){
    fmt.Println("GOが呼び出された")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    var datas = []Data1{}
    var data1 = Data1{}
    var data2 = Data1{Title: "smaple2", Content: "hello, sample2"}
    // fmt.Println(datas)

    data1.Title = "sample1"
    data1.Content = "hello, sample1"
    datas = append(datas, data1)
    datas = append(datas, data2)
    // fmt.Println(datas)
    // jsonエンコード
    outputJson, err := json.Marshal(datas)
    if err != nil {
        panic(err)
    }
    // jsonヘッダーを出力
    w.Header().Set("Content-Type", "application/json")
    fmt.Println(outputJson)
    // jsonデータを出力
    fmt.Fprint(w, string(outputJson))
}