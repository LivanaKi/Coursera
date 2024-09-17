package main

import "fmt"

func factorial(n int64) int64{
	if n == 1{
		return 1
	}else{
		return (n * factorial(n -1))
	}
}

func main(){
	var n int64
	fmt.Println("Enter n positive integer and compute n!")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("n! is %d \n\n", factorial(n))
}