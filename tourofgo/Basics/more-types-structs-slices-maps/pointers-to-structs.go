/*
Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
 */
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := &Vertex{1,2}
	fmt.Println((*p).Y)
	fmt.Println(p.Y)

	v := Vertex{X:3, Y:4}
	fmt.Println(v)
	p = &v
	fmt.Println(p.Y)
	fmt.Println((*p).Y)
}