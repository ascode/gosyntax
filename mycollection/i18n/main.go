package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	message.SetString(language.Chinese, "%s went to %s.", "%s去了%s。")
	message.SetString(language.AmericanEnglish, "%s went to %s.", "%s is in %s.")

	message.SetString(language.Chinese, "%s has been stolen.", "%s被偷走了。")
	message.SetString(language.AmericanEnglish, "%s has been stolen.", "%s has been stolen.")

	message.SetString(language.Chinese, "HOW_ARE_U", "%s 你好吗?")
	message.SetString(language.AmericanEnglish, "HOW_ARE_U", "%s How are you?")
}

func main() {
	//p := message.NewPrinter(language.BritishEnglish)
	//p.Printf("There are %v flowers in our garden.", 1500)
	//
	//fmt.Println("国家语言简写格式-------------------")
	//
	////语言类型构建
	//zh, _ := language.ParseBase("zh") // 语言
	//CN, _ := language.ParseRegion("CN") // 地区
	//zhLngTag, _ := language.Compose(zh, CN)
	//fmt.Println(zhLngTag) // 打印 zh-CN
	//fmt.Println(language.Chinese)// 打印中文缩写
	//fmt.Println(language.SimplifiedChinese)// 打印中文缩写
	//fmt.Println(language.TraditionalChinese)// 打印中文缩写
	//fmt.Println(language.AmericanEnglish)// 打印英文缩写
	//
	//fmt.Println(language.Compose(language.ParseRegion("AL")))
	//// 打印 Und-AL


	// 中文版
	p := message.NewPrinter(language.Chinese)
	p.Printf("%s went to %s.", "彼得", "英格兰")
	fmt.Println()
	p.Printf("%s has been stolen.", "宝石")
	fmt.Println()
	p.Printf("HOW_ARE_U", "竹子")
	fmt.Println()


	// 英文版本
	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("%s went to %s.", "Peter", "England")
	fmt.Println()
	p.Printf("%s has been stolen.", "The Gem")
	fmt.Println()
	p.Printf("HOW_ARE_U", "bamboo")
	fmt.Println()
}
