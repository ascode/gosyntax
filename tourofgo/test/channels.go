package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3}
	c := make(chan int)
	go sum(s, c)
	x := <- c
	fmt.Println(x)


}