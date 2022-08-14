package main

import (
	"fmt"
	"net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    // "reflect"
)
// 構造体名を大文字にしないと以下のエラーになる
// 「struct field title has json tag but is not exportedstructtag」
type Data1 struct {
    gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// グローバルスコープとして定義することで、本ファイルのどの関数でも引数の受け渡しなしに使用可能にする。
var db *gorm.DB
var db_err error

func main() {
    fmt.Println("Start!");
    dsn := "root:@tcp(127.0.0.1:3306)/go_plactice?charset=utf8mb4&parseTime=True&loc=Local"
    db, db_err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if db_err != nil {
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