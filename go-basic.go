package main

import (
	"fmt"
	"math/rand"
    "math"
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
	
}


