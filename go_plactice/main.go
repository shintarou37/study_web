package main

import (
	"fmt"
	"net/http"
    "encoding/json"
)
type Data1 struct {
	Title    string `json:"title"`
	Message  string `json:"content"`
}

func main() {
    fmt.Println("Hello!");
    http.HandleFunc("/", handler);
    http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
    fmt.Println("handler!");
    var data1 = Data1{}
    data1.Title = "sample1"
    data1.Message = "hello, sample1"

    // jsonエンコード
    outputJson, err := json.Marshal(&data1)
    if err != nil {
        panic(err)
    }

    // jsonヘッダーを出力
    w.Header().Set("Content-Type", "application/json")

    // jsonデータを出力
    fmt.Fprint(w, string(outputJson))
}
// test