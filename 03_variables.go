package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // multiple var declaration
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // no value will zero value, for int it's 0
	fmt.Println(e)

	f := "apple" // shorthand for declaration + initialization (f string = "apple")
	fmt.Println(f)
}
