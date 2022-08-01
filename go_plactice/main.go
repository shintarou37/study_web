package main

import (
	"fmt"
	"net/http"
    "encoding/json"
)
type Data1 struct {
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
    fmt.Println("GOが呼び出された")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    var datas = []Data1{}
    var data1 = Data1{}
    var data2 = Data1{"smaple2", "hello, sample2"}
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