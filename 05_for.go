package main

import "fmt"

func main() {

	i := 1 //cazzetto operator, as seen in 03: declaration + init

	for i <= 3 {
		fmt.Println(i)
		i = i + 1 //my shortcut
	}

	for j := 1; j <= 10; j++ { //very C like
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	} //infinite loop

	for n := 0; n < 20; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	} //skipping even

}
