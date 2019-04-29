package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Join("a", "b")
	fmt.Println(p)
}