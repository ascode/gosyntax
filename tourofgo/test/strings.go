package main

import (
	"fmt"
	"strconv"
	"strings"
)

var p = fmt.Println

func main() {
	n := 2342341230000
	p("n:", n)
	p("string(n):", string(n))

	p(strconv.FormatFloat(float64(n), 'f', 0, 64))



	b := 87
	p(strconv.Itoa(b))
	tt := "a" + fmt.Sprintf("%v", 2374897329870000)
	p(tt)

	//fmt.Stringer()

	s1 := "jinfei ssss"
	p(s1[1: 2])


}
