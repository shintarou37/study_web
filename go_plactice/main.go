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
		panic(db_err)
	}

    http.HandleFunc("/", top);
    http.HandleFunc("/detail", detail);
    http.HandleFunc("/edit", edit);
    http.HandleFunc("/register", register);
    http.HandleFunc("/delete", delete);
    http.ListenAndServe(":8080", nil)
    fmt.Println("End!");
}
/* 
    Top画面 
*/
func top(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/\"）でGOが呼び出された")
    // fmt.Fprint(w, "false")
    // 全レコードを取得する
    ret, grm_err := ReadMulti()
    fmt.Println(grm_err)
    if !grm_err{fmt.Println("falseだお")}
    // // jsonエンコード
    outputJson, err := json.Marshal(ret)
    // // エラーが起きた場合はこれ以下の処理は行われない
    if err != nil || !grm_err{
        fmt.Fprint(w, false)
        // panic("err")
    }

    // // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    
    // // jsonデータを返却する
    fmt.Fprint(w, string(outputJson))
}

/* 
    詳細画面
*/
func detail(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/detail\"）でGOが呼び出された")
    var id string = r.URL.Query().Get("id")

    // React側で画面をリロードするとクエリパラメータがundefinedで送付される
    // その場合は"false"という文字列がパラメーターとして送信されてsqlは発行しない
	if id == "false" {
        fmt.Println("error")
		panic("no query params")
	}

    ret, error := Read(id)

    // jsonエンコード
    outputJson, err := json.Marshal(ret)
    if err != nil || error == false{
        panic(err)
    }

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    // jsonデータを返却する
    fmt.Fprint(w, string(outputJson))
}
/* 
    登録機能
*/
func register(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/register\"）でGOが呼び出された")

    // クエリパラメータに含まれた値を使用して構造体を初期化する。
    var create = Data1{Title: r.URL.Query().Get("title"), Content: r.URL.Query().Get("content")}

    // レコードの作成
    if err := db.Create(&create).Error; err != nil {
        fmt.Println("error happen!")
		fmt.Println(err)
        // エラーが起きた場合はエラーページに遷移する
		panic(err)
	}
    
    // jsonエンコード
    outputJson, err := json.Marshal(create)
    if err != nil {
        panic(err)
    }

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    // jsonデータを返却する
    fmt.Fprint(w, string(outputJson))
}

/* 
    編集機能
*/
func edit(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/edit\"）でGOが呼び出された")

    // クエリパラメータに含まれたレコードのIDカラムを取得する
    id := r.URL.Query().Get("id")

    // レコードの更新
    db.Debug().Model(&Data1{}).Where("id = ?", id).Updates(Data1{Title: r.URL.Query().Get("title"), Content: r.URL.Query().Get("content")})
    ret, errr := Read(id)
    fmt.Println(errr)
    // jsonエンコード
    outputJson, err := json.Marshal(ret)
    if err != nil {
        panic(err)
    }

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    // jsonデータを返却する
    fmt.Fprint(w, string(outputJson))
}

/* 
    削除機能
*/
func delete(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/delete\"）でGOが呼び出された")

    // クエリパラメータに含まれたレコードのIDカラムを取得する
    id := r.URL.Query().Get("id")

    // データの削除
    var data1 Data1
    db.Debug().Delete(&data1, id)

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
    // データを返却する
    fmt.Fprint(w, true)
}
/* 
    戻り値を指定していないと
    「too many return values have ([]Data1) want ()compilerWrongResultCount」
    というエラーになる。
    パス：top
*/
func ReadMulti()([]Data1, bool){
    var data1_arr []Data1
    // return data1_arr, false
    if err := db.Debug().Find(&data1_arr).Error; err != nil {
		return data1_arr, false
	}
    return data1_arr, true
}

func Creat(title, content string) bool {
    if err := db.Debug().Create(&Data1{Title: title, Content: content}).Error; err != nil {
        fmt.Println("error happen!")
		fmt.Println(err)
		return false
	}

    return true
}
func CreatMulti(title, content string){
    // multi
    var multi_create = []Data1{{Title: "title2", Content: "content2"}, {Title: "title3", Content: "content3"}, {Title: "title4", Content: "content3"}}
    db.Debug().Create(&multi_create)
    if err := db.Create(&Data1{Title: "title1", Content: "content1"}).Error; err != nil {
        fmt.Println("error happen!")
		fmt.Println(err)
		return
	}
}

func Read(id string) (Data1, bool){
    var data1 Data1
    // ポインタを引数にしない場合はエラーになる
    if err := db.Debug().First(&data1, id).Error; err != nil {
        fmt.Println("error happen!")
		return data1, true
	}
    return data1, true
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