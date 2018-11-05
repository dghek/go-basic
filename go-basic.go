package main

//   Note that: := is a declaration, whereas = is an assignment

//bool
//
//string
//
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//
//byte // alias for uint8
//
//rune // alias for int32
//     
//
//float32 float64
//
//complex64 complex128

// pointers
//  var p *int
//  x = *p

import (
	"fmt"
	"math/rand"
    "math"
	"math/cmplx"
)

var i, j int = 1, 2


func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

//A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
const Pi = 3.14


func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	
    fmt.Println(add(42, 13))
	
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	fmt.Println(split(17))
	
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bo, s)
	
	
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)
	
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

}


