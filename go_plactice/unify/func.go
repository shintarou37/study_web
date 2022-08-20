package unify

import (
	"gorm.io/gorm"
)

type Data1 struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

/*
   パス：top
*/
func ReadMulti(db *gorm.DB) ([]Data1, bool) {
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
func Read(id string, db *gorm.DB) (Data1, bool) {
	var data1 Data1
	// ポインタを引数にしない場合はエラーになる
	if err := db.Debug().First(&data1, id).Error; err != nil {
		return data1, false
	}
	return data1, true
}
