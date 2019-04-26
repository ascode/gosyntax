package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	consul_dc := os.Getenv("CONSUL_DC")
	os.Getenv("")
	fmt.Println("consul_dc:  ", consul_dc)
	envstr, isFound := syscall.Getenv("GOPATH")
	fmt.Printf("%v,%v", envstr, isFound)
}
