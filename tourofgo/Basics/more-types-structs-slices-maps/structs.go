/*
Structs
A struct is a collection of fields.
Struct fields are accessed using a dot.
 */
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{
		X:1,
		Y:2,
	}

	fmt.Println(v.Y)
	v.Y = 1000
	fmt.Println(v.Y)
	fmt.Println(v)
	fmt.Println(Vertex{3,4})
	fmt.Println(Vertex{5,6}.X)

}