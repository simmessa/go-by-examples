package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//multiple case

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend!")
	default: //optional, but always fun, case has no leakage, no break required :)
		fmt.Println("work week :(")
	}

	t := time.Now()

	switch { //on nothing >>> means something similar to an if, please welcome the added flexibility!
	case t.Hour() < 12:
		fmt.Println("AM")
	case t.Hour() >= 12:
		fmt.Println("PM !@")
	}

	//The mystical / mythical beast known as type switch!

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("freaking bool")
		case int:
			fmt.Println("omg INTEGER")
		default:
			fmt.Printf("WTF, a monster: %T\n", t)
		}
	} // wow, pretty complicated shit

	whatAmI(false)
	whatAmI(4)
	whatAmI("prot")
}
