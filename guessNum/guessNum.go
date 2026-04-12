package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
	"math/rand/v2"
)

func main() {
	targetNum := (rand.IntN(100))
	fmt.Println("Start guessing my number!")
	var attempts int = 0
	for {
		attempts++
		fmt.Print("Type a number: ")
		var i int

		fmt.Scan((&i))
		if i == targetNum {
			fmt.Printf("Well Done! It took %v attempts to guess this number.", attempts)
			break
		} else if i < targetNum {
			fmt.Println("My number is greater than", i)
		} else {
			fmt.Println("My number is less than", i)
		}
	}
}
