package main

import (
	"fmt"
	"math/rand"

	"github.com/xantrix/go-basics/stringutil"
)

func add(x, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var i, multi string
	multi = `line
    multi string with Carriage return
    no meaning for \n \r
    `
	i = "string " + "concatenation " + multi

	fmt.Printf("Hello000, world \n" + i)
	fmt.Println("Reverse:" + stringutil.Reverse("!oG ,olleH"))

	fmt.Println("My favorite number is", rand.Intn(10))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
