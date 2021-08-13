package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// const n = 234 // can I redefine / redeclare this? NOPE

	const d = 3e20 / n //wow exp notation

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
