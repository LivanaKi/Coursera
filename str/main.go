package main

import (
	"fmt"
)

func main(){
	var c byte = 'a'
	var strOne string = "My Idea?"

	fmt.Println("String Program")
	fmt.Println("first is:", strOne)
	fmt.Println("c is:", c)

	sum := len(strOne)
	for i := 0; i < sum; i++{
		fmt.Printf("char %d is %c\n", i, strOne[i])
	}
}