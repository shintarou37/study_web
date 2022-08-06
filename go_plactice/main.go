package main

import (
	"fmt"
	"net/http"
    // "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "reflect"
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
    http.HandleFunc("/", handler);
    http.ListenAndServe(":8080", nil)
    fmt.Println("End!");
}

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/\"）でGOが呼び出された")
    dsn := "root:@tcp(127.0.0.1:3306)/go_plactice?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		panic("failed to connect database")
	}
    ret := ReadMulti(db)
    fmt.Println("戻り値を出力する")
    fmt.Println(ret)

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

    // // ヘッダーをセットする
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    // w.Header().Set("Content-Type", "application/json")
    // // jsonを出力
    // fmt.Println(outputJson)
    // // jsonデータを出力
    // fmt.Fprint(w, string(outputJson))
}

func Creat(db *gorm.DB){
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
func ReadMulti(db *gorm.DB)[]Data1{
    var data1_arr []Data1
    db.Debug().Find(&data1_arr)
    fmt.Println(data1_arr)
    fmt.Println(reflect.TypeOf(data1_arr))
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