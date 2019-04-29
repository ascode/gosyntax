/*
我的理解是任何结构体都实现了一个空的接口interface{},所以使用interface{}类型就能够接受任何结构体
 */
package interface_and_type_switch

import "fmt"

var omnipotent interface{}

type imen interface {
	Aa()
}

type men struct {
	name string
	age int
}

func (m men)Aa() {
	fmt.Println("this is a man")
}

func receiveAObject (cfg interface{}) interface{} {
	fmt.Println("begin.....")
	_ = 1
	//fmt.Println(cfg.(type)) // 编译报错 use of .(type) outside type switch
	switch omnipotent := cfg.(type) {
	case *men: // TODO: 这个地方*men在内存呢中做了什么吗
		fmt.Println("a omnipotent type is: ", *omnipotent)
	case men:
		fmt.Println("omnipotent type is: ", omnipotent)
	default:
		fmt.Println("不是人")
	}
	return cfg
}

func main() {
	fmt.Println(receiveAObject(1))
	fmt.Println(receiveAObject("jinfei"))
	fmt.Println(receiveAObject(men{
		name: "jinfei",
		age: 30,
	}))

	fmt.Println("---------------------------")

	//aMen := &men{}
	//aMen.Aa()
	//switch a := aMen.(type) {
	//case men:
	//	fmt.Println(a)
	//}

	var bMen imen
	bMen = men{}
	bMen.Aa()

	switch a := bMen.(type) {
	case men:
		fmt.Println(a)
	}
}

