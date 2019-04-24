package main

import "fmt"

func main() {
	s := []int {1, 2, 1, 3, 2, 5, 6}
	fmt.Println(Filter(s, func(i int) bool {
		if i == 5 {
			return true
		} else {
			return false
		}
	}))
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}