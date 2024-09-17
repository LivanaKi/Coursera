package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var pair, vPair, howMany int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("reading pair of dice value")
	fmt.Scanf("%d", &vPair)
	for i := 0; i < 10000000; i++{
		pair = rand.Intn(6) + 2
		if pair == vPair{
			howMany++
		}
	}
	fmt.Println(howMany)
}