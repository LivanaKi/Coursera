package main

import(
	"fmt"
)

var myInt = 5

func main(){
	fmt.Println("Will try some simple ideas ")
	fmt.Println("Print the global", myInt)
	{
		var myInt int = 6
		fmt.Println("Print the inner", myInt)
	}
	fmt.Println("Print thr global", myInt)
}