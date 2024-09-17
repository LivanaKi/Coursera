/*
Program that use random number generation to solve probability of 2 (or more) having the same birthday.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sameBirthday(countPeople int, countTrials int) float64 {

	var countDublicate int

	for i := 0; i < countTrials; i++ {
		birthdays := make(map[int]bool) //save data and check duplicate
		duplicate := false

		for j := 0; j < countPeople; j++ {
			dayBirthday := rand.Intn(365) + 1 //generation day

			if birthdays[dayBirthday] { // if birthdays match, loop break
				duplicate = true
				break
			}
			birthdays[dayBirthday] = true
		}
		if duplicate { // if true, countDublicate plus one
			countDublicate++
		}
		//fmt.Println(birthdays)
	}
	return float64(countDublicate) / float64(countTrials) // probability
}

func main() {
	fmt.Println("Program that use random number generation to solve probability of 2 (or more) having the same birthday")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())
	fmt.Println("People\tProbability")
	trials := 10000

	for amountPeople := 10; amountPeople <= 100; amountPeople += 10 {
		probability := sameBirthday(amountPeople, trials)
		fmt.Printf("%d\t%.4f\n", amountPeople, probability)
	}
}
