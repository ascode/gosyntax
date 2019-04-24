package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
	//"test/add1"
)

/*The example shows variables of several types, and also that variable
declarations may be "factored" into blocks, as with import statements.
*/
var (
	ToBe bool		  = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func main(){
	fmt.Println("Hello,world!")
	fmt.Println(time.Now())
	fmt.Println(rand.Int63n(time.Now().Unix()))
	fmt.Println(math.Pi)
	fmt.Println(add(3,4))
	//fmt.Println(add1)
	fmt.Println(processAToAB(2))
	fmt.Println("c = ", c)
	fmt.Println("i + j = ",i + j)

	k := 1000//Short variable declarations :=短变量声明
	fmt.Println(k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


	// Zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println(Big * 0.1 / 0.1)
	fmt.Println(float64(Big))
	fmt.Println(Small)
	fmt.Println("-------------")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
