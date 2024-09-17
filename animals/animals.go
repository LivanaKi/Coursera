package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat string
	move string
	speak  string
}

func (animal *Animal) Eat() {
	fmt.Println("Eat:", animal.eat)
}
func (animal *Animal) Move() {
	fmt.Println("Move:", animal.move)
}
func (animal *Animal) Speak() {
	fmt.Println("Sound:", animal.speak)
}

func animalFunc(animal Animal, line1 string) {
	switch line1 {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("There is no such information")

	}
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	fmt.Println("Enter name animals and another inf")
	fmt.Println("If you want to complete the program, enter -> ok :)")
	for {
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line1 := strings.Split(line, " ")

	if line1[0] == "cow" {
		animals := cow
		animalFunc(animals, line1[1])
	} else if line1[0] == "bird" {
		animals := bird
		animalFunc(animals, line1[1])
	} else if line1[0] == "snake" {
		animals := snake
		animalFunc(animals, line1[1])
	} else if line1[0] == "ok"{
		fmt.Println("Stop!")
		break
	}else {
		fmt.Println("there is no such animal")
	}
}
}
