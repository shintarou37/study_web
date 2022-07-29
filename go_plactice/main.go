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
    var data1 = Data1{}
    // var data2 = Data1{"2title","2content"}
    data1.Title = "sample1"
    data1.Content = "hello, sample1"

    // jsonエンコード
    outputJson, err := json.Marshal(&data1)
    // outputJson2, err2 := json.Marshal(&data2)
    if err != nil /*|| err2 != nil*/ {
        panic(err)
    }
    // fmt.Println(string(outputJson));
    // fmt.Println(string(outputJson2));
    // jsonヘッダーを出力
    w.Header().Set("Content-Type", "application/json")

    // jsonデータを出力
    fmt.Fprint(w, string(outputJson))
}