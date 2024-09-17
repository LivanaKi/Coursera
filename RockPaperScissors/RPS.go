// Go program for game Rock-Paper-Scissors

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rock     = 0
	paper    = 1
	scissors = 2
)
const (
	cRock     = 'R'
	cPaper    = 'P'
	cScissors = 'S'
)

func translate(move int) rune {
	var cMove rune
	if move == rock {
		cMove = cRock
	} else if move == paper {
		cMove = cPaper
	} else if move == scissors {
		cMove = cScissors
	}
	return cMove
}

func main() {
	var move, machineMove int
	var count int

	var draws, wins, machineWins int
	var data map[int]int
	data = make(map[int]int)
	rounds := 100

	fmt.Print("-----START 100 ROUNDS----")

	for i := 0; i < rounds; i++ {
		//User plays

		fmt.Println("\nRound ", i+1)
		rand.Seed(time.Now().UnixNano())
		move = rand.Intn(3)
		count = 20

		if i < 10 {
			machineMove = rand.Intn(3)
		} else if i > 10 && i < count {
			var choosePlay int
			for key, val := range data {
				max := 0
				if max < val {
					max = val
				}
				if data[key] == max {
					choosePlay = key
				}
			}
			switch choosePlay {
			case rock:
				machineMove = paper
			case paper:
				machineMove = scissors
			case scissors:
				machineMove = rock
			}
		} else {
			clear(data)
			count += 10
		}
		data[move]++
		switch move {
		case rock:
			if machineMove == rock {
				draws++
			} else if machineMove == paper {
				machineWins++
			} else {
				wins++
			}
		case paper:
			if machineMove == paper {
				draws++
			} else if machineMove == scissors {
				machineWins++
			} else {
				wins++
			}
		case scissors:
			if machineMove == scissors {
				draws++
			} else if machineMove == rock {
				machineWins++
			} else {
				wins++
			}
		}
		fmt.Printf("Player plays: %c\n", translate(move))
		fmt.Printf("Machine plays: %c\n", translate(machineMove))

	}
	fmt.Println("\n After ", rounds, "rounds:\n", "you win:", wins, ", machine wins", machineWins, "with", draws, "draws")
}
