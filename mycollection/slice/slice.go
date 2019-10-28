package main

import "fmt"

func main() {
	aa := []int{1,2,3,4}
	test1(aa)
	fmt.Println(aa)
	tt := aa[:2]
	fmt.Println("切片修改前", cap(tt))
	test1(tt)
	fmt.Println("切片修改后在方法外", cap(tt))
	fmt.Println(aa)
}

func test1(a []int) {
	a = append(a, 123)
	a[0] = 888
	fmt.Println("切片在方法中修改后", a)
}