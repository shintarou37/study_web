package main

import (
	"fmt"
	"net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    // "reflect"
)
const (
	StatusInternalServerError = 500
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

    // ヘッダーをセットする（エラー処理後にセットするとCROSエラーになる）
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Content-Type", "application/json")

    // 全レコードを取得する
    ret, orm_err := ReadMulti()

    // jsonエンコード
    outputJson, err := json.Marshal(ret)

    // エラー処理
    if err != nil || !orm_err{
        fmt.Println("error happen!")
        w.WriteHeader(http.StatusInternalServerError)
    }

    // jsonデータを返却する（エラーが発生した場合は空のオブジェクトを返却する）
    fmt.Fprint(w, string(outputJson))
}

/* 
    詳細画面
*/
func detail(w http.ResponseWriter, r *http.Request){
    fmt.Println("パス（\"/detail\"）でGOが呼び出された")

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    // クエリパラメータ「id」を取得する
    var id string = r.URL.Query().Get("id")

    // React側で画面をリロードするとクエリパラメータがundefinedで送付される
    // その場合は"false"という文字列がパラメーターとして送信されてsqlは発行しない
	if id == "false" {
		panic("no query params")
        // これ以降の処理は行われない
	}
    
    ret, orm_err := Read(id)

    // jsonエンコード
    outputJson, err := json.Marshal(ret)

    // エラー処理
    if err != nil || !orm_err {
        fmt.Println("error happen!")
        w.WriteHeader(http.StatusInternalServerError)
    }

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
    if orm_err := db.Create(&create).Error; orm_err != nil {
        fmt.Println("error happen!")
		w.WriteHeader(http.StatusInternalServerError)
	}
    
    // jsonエンコード
    outputJson, err := json.Marshal(create)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
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
    if orm_err := db.Debug().Model(&Data1{}).Where("id = ?", id).Updates(Data1{Title: r.URL.Query().Get("title"), Content: r.URL.Query().Get("content")}).Error; orm_err != nil {
        fmt.Println("error happen!")
		w.WriteHeader(http.StatusInternalServerError)
	}

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    // 任意のデータを返却する（データは使用しないので値は任意）
    fmt.Fprint(w, string(id))
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
    if orm_err := db.Debug().Delete(&data1, id).Error; orm_err != nil {
        fmt.Println("error happen!")
		w.WriteHeader(http.StatusInternalServerError)
	}

    // ヘッダーをセットする
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Methods","GET, DELETE")

    // 任意のデータを返却する（データは使用しないので値は任意）
    fmt.Fprint(w, string(id))
}
/* 
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

/* 
    パス：detail
*/
func Read(id string) (Data1, bool){
    var data1 Data1
    // return data1, false
    // ポインタを引数にしない場合はエラーになる
    if err := db.Debug().First(&data1, id).Error; err != nil {
        fmt.Println("error happen!")
		return data1, true
	}
    return data1, true
}