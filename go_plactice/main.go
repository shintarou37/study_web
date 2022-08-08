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

    http.HandleFunc("/", top);
    http.HandleFunc("/register", register);
    http.ListenAndServe(":8080", nil)
    fmt.Println("End!");
}

func top(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/\"）でGOが呼び出された")
    ret := ReadMulti()

    // var datas = []Data1{}
    // var data1 = Data1{}
    // var data2 = Data1{Title: "smaple2", Content: "hello, sample2"}
    // // fmt.Println(datas)
    // data1.Title = "sample1"
    // data1.Content = "hello, sample1"
    // datas = append(datas, data1)
    // datas = append(datas, data2)
    // // fmt.Println(datas)

    // jsonエンコード
    outputJson, err := json.Marshal(ret)
    if err != nil {
        panic(err)
    }

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    
    // jsonをコンソールに出力する
    // fmt.Println(string(outputJson))
    // jsonデータを返却する
    fmt.Fprint(w, string(outputJson))
}

func register(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/register\"）でGOが呼び出された")
    var data2 = Data1{Title: "smaple2", Content: "hello, sample2"}
    outputJson, err := json.Marshal(data2)
    if err != nil {
        panic(err)
    }

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    fmt.Fprint(w, string(outputJson))
}

func Creat(){
    db.Debug().Create(&Data1{Title: "title1", Content: "content1"})
    // multi
    var multi_create = []Data1{{Title: "title2", Content: "content2"}, {Title: "title3", Content: "content3"}, {Title: "title4", Content: "content3"}}
    db.Debug().Create(&multi_create)
    if err := db.Create(&Data1{Title: "title1", Content: "content1"}).Error; err != nil {
        fmt.Println("error happen!")
		fmt.Println(err)
		return
	}
}
/* 
    戻り値を指定していないと
    「too many return values have ([]Data1) want ()compilerWrongResultCount」
    というエラーになる。
*/
func ReadMulti()[]Data1{
    var data1_arr []Data1
    db.Debug().Find(&data1_arr)
    return data1_arr
}

func Read(db *gorm.DB){
    var data1 Data1
    db.Debug().First(data1, 2)
    fmt.Println(data1)
}

func Update(db *gorm.DB){
    var data1 Data1
    if err := db.Debug().Model(&Data1{}).Where("id = ?", 1).Update("Title", "titleUpdate").Error; err != nil {
		fmt.Println(err)
		fmt.Println(data1)
		return
	}

    // multi
    db.Debug().Model(&data1).Where("id = ?", 2).Updates(Data1{Title: "titleMultUpdate", Content: "ContentMultUpdate"})
}

func Delete(db *gorm.DB){
    var data1 Data1
    db.Debug().Delete(&data1, 1)
}