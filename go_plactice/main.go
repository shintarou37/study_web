package main

import (
	"fmt"
	"net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)
// 構造体名を大文字にしないと以下のエラーになる
// 「struct field title has json tag but is not exportedstructtag」
type Data1 struct {
    gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func main() {
    fmt.Println("Start!");
    // dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    dsn := "root:secualpass@tcp(127.0.0.1:3306)/go_plactice?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		panic("failed to connect database")
	}
    var data1 Data1
    // Create
    // db.Create(&Data1{Title: "title1", Content: "content1"})

    // Read
    db.Find(&data1)
    fmt.Println(data1)
    // {{1 2022-08-04 09:52:35.172 +0900 JST 2022-08-04 12:30:43.891 +0900 JST {0001-01-01 00:00:00 +0000 UTC false}} titleMultUpdate ContentMultUpdate}
    outputJson, err := json.Marshal(data1)
    fmt.Println(string(outputJson))
    // {"ID":1,"CreatedAt":"2022-08-04T09:52:35.172+09:00","UpdatedAt":"2022-08-04T12:30:43.891+09:00","DeletedAt":null,"title":"titleMultUpdate","content":"ContentMultUpdate"}

    // Update
    // 「db.Model(&data1).Update("title1", "titleUpdate")」だとエラーになる。以下URLより、Update文ではwhere句を使用するべきである。エラー文「WHERE conditions required」
    // https://gorm.io/docs/update.html#Block-Global-Updates
    // if err := db.Model(&Data1{}).Where("id = ?", 1).Update("Title", "titleUpdate").Error; err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(data1)
	// 	return
	// }

    // Update Multi
    // db.Model(&data1).Where("id = ?", 1).Updates(Data1{Title: "titleMultUpdate", Content: "ContentMultUpdate"})

    // Delete
    // if err := db.Delete(&data1, 1).Error; err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(data1)
	// 	return
	// }

    // エラー文「invalid value, should be pointer to struct or slice」
    // if err := db.Delete(data1, 1).Error; err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(data1)
	// 	return
	// }

    http.HandleFunc("/", handler);
    // http.ListenAndServe(":8080", nil)
    fmt.Println("End!");

}
func handler(w http.ResponseWriter, r *http.Request){
    // fmt.Println("GOが呼び出された")
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    // var datas = []Data1{}
    // var data1 = Data1{}
    // var data2 = Data1{Title: "smaple2", Content: "hello, sample2"}
    // // fmt.Println(datas)

    // data1.Title = "sample1"
    // data1.Content = "hello, sample1"
    // datas = append(datas, data1)
    // datas = append(datas, data2)
    // // fmt.Println(datas)
    // // jsonエンコード
    // outputJson, err := json.Marshal(datas)
    // if err != nil {
    //     panic(err)
    // }
    // // jsonヘッダーを出力
    // w.Header().Set("Content-Type", "application/json")
    // fmt.Println(outputJson)
    // // jsonデータを出力
    // fmt.Fprint(w, string(outputJson))
}