package models

import (
	"math/rand"
	"time"
)

type Englishwords struct {
	Id      int `gorm:type:Auto_Increment,not null`
	Words   string
	Chapter int
	Unit    int
	Count   int
}

// 数据库中获得所有
func GetAllWords() ([]*Englishwords, error) {
	var allwords []*Englishwords
	result := db.Find(&allwords)
	if result.Error != nil {
		return nil, result.Error
	}
	return Randresult(allwords), nil
}

// 通过章节获取
func GetWordsByChapter(chapter int) ([]*Englishwords, error) {
	var allwords []*Englishwords
	result := db.Where("chapter = ?", chapter).Find(&allwords)
	if result.Error != nil {
		return nil, result.Error
	}
	return allwords, nil
}

// 通过综合查询获取
func GetWords(inquery map[string]interface{}) ([]*Englishwords, error) {
	var patternwords []*Englishwords
	result := db.Where(inquery).Find(&patternwords)
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
