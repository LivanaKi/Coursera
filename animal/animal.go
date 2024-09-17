package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name  []string
	eat   string
	move  string
	speak string
}

func (animal *Cow) Eat() {
	fmt.Println("Eat:", animal.eat)
}
func (animal *Cow) Move() {
	fmt.Println("Move:", animal.move)
}
func (animal *Cow) Speak() {
	fmt.Println("Sound:", animal.speak)
}

type Bird struct {
	name  []string
	eat   string
	move  string
	speak string
}

func (animal *Bird) Eat() {
	fmt.Println("Eat:", animal.eat)
}
func (animal *Bird) Move() {
	fmt.Println("Move:", animal.move)
}
func (animal *Bird) Speak() {
	fmt.Println("Sound:", animal.speak)
}

type Snake struct {
	name  []string
	eat   string
	move  string
	speak string
}

func (animal *Snake) Eat() {
	fmt.Println("Eat:", animal.eat)
}
func (animal *Snake) Move() {
	fmt.Println("Move:", animal.move)
}
func (animal *Snake) Speak() {
	fmt.Println("Sound:", animal.speak)
}

func nameAnimal(line string, name []string) bool {
	for _, val := range name {
		if line == val {
			return true
		}
	}
	return false
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
	cowNew := Cow{nil, "grass", "walk", "moo"}
	birdNew := Bird{nil, "worms", "fly", "peep"}
	snakeNew := Snake{nil, "mice", "slither", "hsss"}

	var animal Animal

	fmt.Println("Enter newanimal/query, name animals and\ntype (cow, bird, snack)/information(eat, move, speak)")
	fmt.Println("If you want to complete the program, enter -> ok :)")
	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		line1 := strings.Split(line, " ")

		if line1[0] == "newanimal" {
			if line1[2] == "cow" {
				cowNew.name = append(cowNew.name, line1[1])
				fmt.Println("Created it!")
			} else if line1[2] == "bird" {
				birdNew.name = append(birdNew.name, line1[1])
				fmt.Println("Created it!")
			} else if line1[2] == "snake" {
				snakeNew.name = append(snakeNew.name, line1[1])
				fmt.Println("Created it!")
			} else {
				fmt.Println("No such type exists")
			}
		} else if line1[0] == "query" {
			if nameAnimal(line1[1], cowNew.name) {
				animal = &cowNew
				animalFunc(animal, line1[2])
			} else if nameAnimal(line1[1], birdNew.name) {
				animal = &birdNew
				animalFunc(animal, line1[2])
			} else if nameAnimal(line1[1], snakeNew.name) {
				animal = &snakeNew
				animalFunc(animal, line1[2])
			} else {
				fmt.Println("The name was entered incorrectly or does not exist")
			}
		} else if line1[0] == "ok" {
			fmt.Println("Stop!")
			break
		} else {
			fmt.Println("The command was entered incorrectly")
		}
	}
}
