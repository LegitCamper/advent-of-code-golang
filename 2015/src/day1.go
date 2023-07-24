package main

import "fmt"

func main() {
	var part2 = true
	var floors string = ""

	var floor int = 0
	for i := 0; i < len(floors); i++ {
		if string(floors[i]) == "(" {
			floor++
		} else if string(floors[i]) == ")" {
			floor--
		}
		if part2 && floor < 0 {
			fmt.Print("Santa went into the basement at: ", i+1, "\n")
			break
		}
	}

	fmt.Print("Santas floor is: ", floor)
}
