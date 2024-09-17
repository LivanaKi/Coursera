package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var pair, howMan int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000000; i++{
		pair = rand.Intn(6) + rand.Intn(6) + 2
		if pair == 7{
			howMan++
		}
	}
	fmt.Println(howMan)
	fmt.Printf("Probability of rolling a seven is %4g \n", float32(howMan)/10000000.0)
}