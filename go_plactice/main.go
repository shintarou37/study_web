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
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	StatusFound            = 302 // RFC 9110, 15.4.3
	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
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
    ret, grm_err := ReadMulti()

    // jsonエンコード
    outputJson, err := json.Marshal(ret)

    // エラー処理
    if err != nil || !grm_err{
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
    
    ret, grm_err := Read(id)

    // jsonエンコード
    outputJson, err := json.Marshal(ret)

    // エラー処理
    if err != nil || !grm_err {
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
    // return data1, false
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