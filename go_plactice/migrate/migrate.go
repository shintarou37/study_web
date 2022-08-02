package main

import (
	"fmt"
    "gorm.io/gorm"
    "01up"
    "01down"
)
type Data1 struct {
    gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func main() {
    fmt.Println("Start migrate!");
    // Up001()
    01down.Down001()
    fmt.Println("End migrate!");
}