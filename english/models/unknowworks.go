package models

import "fmt"

type Unknowwords struct {
	Id     int `gorm:type:Auto_Increment,not null`
	Unknow int `gorm:size:4`
}

func CreateUnKnowWords(unknow int) {
	newUnKownWord := Unknowwords{Unknow: unknow}
	result := db.Create(&newUnKownWord)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	//fmt.Println(newUnKownWord.Id)
}

func DeleteUnKnowWords(unknow int) {
	result := db.Where("unknow = ?", unknow).Delete(&Unknowwords{})
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func DeleteAllWords() {
	result := db.Exec("delete from unknowwords")
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func SelectUnKnowWords() ([]*Englishwords, error) {
	var results []*Englishwords
	//db.Table("englishwords").Select("englishwords.id,englishwords.words,englishwords.chapter," +
	//   "englishwords.unit").Joins("join unknowwords where unknowwords.unknow = englishwords.id").Scan(&results)
	db.Table("englishwords").Select("englishwords.*").Joins("join unknowwords where unknowwords.unknow = englishwords.id").Scan(&results)
	if db.Error != nil {
		return nil, db.Error
	}
	if len(results) == 0 {
		return nil, nil
	}
	return Randresult(results), nil
}
