package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func init() {
	// 以上代码和以下代码都是硬编码方式
	for _, e := range msaArry {
		tag := language.MustParse(e.tag)
		switch msg := e.msg.(type) {
		case string:
			message.SetString(tag, e.key, msg)
		case catalog.Message:
			message.Set(tag, e.key, msg)
		case []catalog.Message:
			message.Set(tag, e.key, msg...)
		}
	}
}

func main() {

	//调用硬编码中的消息 国际化
	p := message.NewPrinter(language.English)
	p.Printf("HELLO_WORLD","bamboo")
	p.Println()
	p.Printf("TASK_REM", 2)
	p.Println()

}

//手工 硬编码方式
var msaArry = [...]langMsg{
	{"en", "HELLO_WORLD", "%s Hello World"},
	{"zh", "HELLO_WORLD", "%s 你好世界"},
	{"en", "TASK_REM", plural.Selectf(1, "%d",
		"=1", "One task remaining!",
		"=2", "Two tasks remaining!",
		"other", "[1]d tasks remaining!",
	)},
	{"zh", "TASK_REM", plural.Selectf(1, "%d",
		"=1", "剩余一项任务！",
		"=2", "剩余两项任务！",
		"other", "剩余 [1]d 项任务！",
	)},
}