package services

import (
	"fmt"
	"mygo/english/models"
	"os"
)

func StartApp() {
	for {
		showMenu()
		var inputNum int
		fmt.Print("请输入需要的序列号: ")
		fmt.Scanf("%d\n", &inputNum)
		//fmt.Printf("用户输入的是： %d \n", inputNum)

		switch inputNum {
		case 1:
			var inputUnit int
			fmt.Print("请输入单元的序列号1-47: ")
			fmt.Scanf("%d\n", &inputUnit)
			SelectGetWords(1, inputUnit)
		case 2:
			var inputChapter int
			fmt.Print("请输入单元的序列号1-10: ")
			fmt.Scanf("%d\n", &inputChapter)
			SelectGetWords(2, inputChapter)
		case 3:
			SelectForgot()
		case 4:
			DeleteForgot()
		case 5:
			SelectAllWords()
		case 6:
			models.DbClose()
			os.Exit(0)
		}
	}

}
func showMenu() {
	fmt.Println("")
	fmt.Println("欢迎来到背单词系统")
	fmt.Println("1 单元单词")
	fmt.Println("2 章节单词")
	fmt.Println("3 遗忘单词")
	fmt.Println("4 删除遗忘")
	fmt.Println("5 全部单词")
	fmt.Println("6 退出系统")
}
func SelectAllWords() {
	words, _ := models.GetAllWords()
	var inputStr string
	sumWords := len(words)
	nextWords := 0
	for i := 0; i < sumWords; i++ {
	gotoHere:
		fmt.Println("")
		fmt.Printf("n:下一个, p:跳过, q:返回, 比例: %d/%d %.2f%% \n", i+1, sumWords, float64(nextWords)/float64(sumWords)*100)
		fmt.Println(words[i].Words)
		fmt.Scanf("%s\n", &inputStr)
		//fmt.Printf("用户输入的是： %d \n", inputStr)
		switch inputStr {
		case "n", "N":
			nextWords = nextWords + 1
			continue
		case "p", "P":
			go models.CreateUnKnowWords(words[i].Id)
		case "q", "Q":
			StartApp()
		default:
			goto gotoHere
		}
	}

}

func DeleteForgot() {
	models.DeleteAllWords()
	fmt.Println("删除成功")
	StartApp()
}

func SelectForgot() {
	words, _ := models.SelectUnKnowWords()
	if words == nil {
		fmt.Println("")
		fmt.Println("没有单词,返回上层!!!")
		StartApp()
	}
	var inputStr string
	sumWords := len(words)
	nextWords := 0
	for i := 0; i < sumWords; i++ {
	gotoHere:
		fmt.Println("")
		fmt.Printf("n:下一个, p:跳过, q:返回, 比例: %d/%d %.2f%% \n", i+1, sumWords, float64(nextWords)/float64(sumWords)*100)
		fmt.Println(words[i].Words)
		fmt.Scanf("%s\n", &inputStr)
		//fmt.Printf("用户输入的是： %s \n", inputStr)
		switch inputStr {
		case "n", "N":
			//fmt.Println(words[i].Id)
			nextWords = nextWords + 1
			go models.DeleteUnKnowWords(words[i].Id)
		case "p", "P":
			continue
		case "q", "Q":
			StartApp()
		default:
			goto gotoHere
		}
	}

}
func SelectGetWords(number int, querynumber int) {
	var inquery = make(map[string]interface{})
	switch number {
	case 1:
		inquery["unit"] = querynumber
	case 2:
		inquery["chapter"] = querynumber
	}

	words, _ := models.GetWords(inquery)
	sumWords := len(words)
	nextWords := 0
	var inputStr string
	for i := 0; i < sumWords; i++ {
	gotoHere:
		fmt.Println("")
		fmt.Printf("n:下一个, p:跳过, q:返回, 比例: %d/%d %.2f%% \n", i+1, sumWords, float64(nextWords)/float64(sumWords)*100)
		fmt.Println(words[i].Words)
		fmt.Scanf("%s\n", &inputStr)
		//fmt.Printf("用户输入的是： %s \n", inputStr)
		switch inputStr {
		case "n", "N":
			nextWords = nextWords + 1
			continue
		case "p", "P":
			go models.CreateUnKnowWords(words[i].Id)
		case "q", "Q":
			StartApp()
		default:
			goto gotoHere
		}
	}
}
