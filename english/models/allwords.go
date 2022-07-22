package models

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Englishwords struct {
	Id      int    `gorm:type:Auto_Increment,not null`
	Words   string `json:"words"`
	Chapter int    `json:"chapter"`
	Unit    int    `json:"unit"`
	TSL     string `json:"tsl"`
	Count   int    `json:count`
}

// 数据库中获得所有
func GetAllWords() ([]*Englishwords, error) {
	var allwords []*Englishwords
	result := db.Find(&allwords).Where("status = 1")
	if result.Error != nil {
		return nil, result.Error
	}
	return Randresult(allwords), nil
}

// 通过章节获取
func GetWordsByChapter(chapter int) ([]*Englishwords, error) {
	var allwords []*Englishwords
	result := db.Where("chapter = ?", chapter).Where("status = 1 ").Find(&allwords)
	if result.Error != nil {
		return nil, result.Error
	}
	return allwords, nil
}

// 通过综合查询获取
func GetWords(inquery map[string]interface{}) ([]*Englishwords, error) {
	var patternwords []*Englishwords
	result := db.Where(inquery).Where("status = 1").Find(&patternwords)
	if result.Error != nil {
		return nil, result.Error
	}
	return Randresult(patternwords), nil

}

// 随机获取结果
func Randresult(words []*Englishwords) []*Englishwords {
	//func Randresult(inquery map[string]interface{})  {
	rand.Seed(time.Now().Unix())
	// 高纳德置乱算法(洗牌算法)
	for i := len(words) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		words[i], words[num] = words[num], words[i]
	}
	return words
}

func AddCount(word *Englishwords) {
	result := db.Model(&word).Where("id = ?", word.Id).Update("count", gorm.Expr("count+ ?", 1))
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
