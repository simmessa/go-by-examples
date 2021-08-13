package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	if 8%4 == 0 {
		fmt.Println("sucks")
	}

	if num := 9; num < 0 {
		fmt.Println("neg")
	} else if num < 10 {
		fmt.Println("1 digit")
	} else if num > 10 {
		fmt.Println("2 digits")
	}

}

//REMEMBER: >>> There is no ternary op in Go. For better or worse.
