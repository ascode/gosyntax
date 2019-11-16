package main

import (
	"fmt"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {

	//根据参数处理不同的返回语句
	message.Set(language.English, "APP_COUNT",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))
}

func main() {
	fmt.Println("placehold中的条件判断-------------------")
	p := message.NewPrinter(language.English)
	// 条件判断
	p.Printf("APP_COUNT", 1)
	fmt.Println()
	p.Printf("APP_COUNT", 2)
	p.Println()
}
