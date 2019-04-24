package main

import "fmt"

func sum2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println(len(s))
	fmt.Println(s[len(s)/2:])
	fmt.Println(s[:3])
	c := make(chan int)
	go sum2(s[:len(s)/2], c)
	go sum2(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}

