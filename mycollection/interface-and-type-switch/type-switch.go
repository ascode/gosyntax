/*
Type switch是个什么鬼
目前我只能这么理解：
1.go里的type switch不是普通的switch,就是go语言在语法层面为我们设计的类型转换的一种模式。
2.这里的t:=必须和switch结合在一起.
 */
package main

import "fmt"

func main() {
	var arg interface{}
	arg = 1
	switch t := arg.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t)             // t has type bool
	case int:
		fmt.Printf("integer %d\n", t)             // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}