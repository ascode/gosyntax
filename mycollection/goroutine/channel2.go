package main

import (
	"fmt"
	"time"
)


var ch = make(chan int)
var j = 0
func say() {
	ch <- 1
	time.Sleep(100000*20000)
	j-= <- ch
}

func main() {

	for j < 10 {
		go say()
		fmt.Println(j)
		j += <- ch

	}
	fmt.Println("l")

	for j < 10 {
		fmt.Println("k")
		fmt.Println(<-ch)
	}


}

