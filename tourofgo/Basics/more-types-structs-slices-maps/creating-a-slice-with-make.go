/*
Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
 */
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
	if b == nil {
		fmt.Println("b value: ", "nil")
	}

	var e []int
	printSlice("e", e)
	if e == nil {
		fmt.Println("e value: ", "nil")
	}

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)


	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println("p;;;", p)

	x1 := []int{11,12,13}
	x2 := []int{14, 15, 16 }
	x3 := append(x1, x2...)
	println("x1---", x1)
	fmt.Println("x3---", x3)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func AppendByte(slice []byte, data ...byte) []byte {
	fmt.Println("data:", []byte(data), data[0], data[1])
	fmt.Println("data[0]", data[0])
	fmt.Println("data[1]", data[1])
	fmt.Println("data[2]", data[2])
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}