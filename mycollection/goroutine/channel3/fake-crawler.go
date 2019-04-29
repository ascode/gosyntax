package main

import (
	"fmt"
	"time"
)

func test(t int) {
	fmt.Println("我在爬数据", t)
	time.Sleep(4000000000)
	c <- t
}
var c = make(chan int)
var ilist []int
func main() {
	for i := 0;i < 13; i++ {
		go test(i)
	}


	for i := 0;i < 13; i++ {
		t := <-c
		fmt.Println("我返回了收到的数据", t)
		ilist = append(ilist, t)
	}

	fmt.Println("我爬到的全部数据：", ilist)
}