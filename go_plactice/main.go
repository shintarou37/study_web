package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("Hello!");
    http.HandleFunc("/", handler);
    http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request){
    fmt.Println("handler!");
}