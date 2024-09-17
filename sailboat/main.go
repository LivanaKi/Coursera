package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 15; j++ {
			if j == 7-i && i < 6 {
				fmt.Print("*")
			} else if i == 6 && j < 8 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			if j == 7 && i < 8 { 
				fmt.Print("|")
			}
		}
		fmt.Println()
	}

	b := 3
	for i := 0; i < 15; i++ {
		fmt.Print("#")
	}
	fmt.Println()
	for i := 0; i < 1; i++ {
		fmt.Print("#")
		for j := 0; j < 13; j++ {
			fmt.Print(" ")
		}
		fmt.Println("#")
	}
	if b > 1 {
		i := 0
		for i < 15 {
			if i == 0 || i == 14 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
			i++
		}
		fmt.Println()
	}
}
