package main

import(
	"fmt"
)

func main(){
	var area, radius float32
	const pi = 3.14159

	fmt.Println("Input radius")
	fmt.Scanf("%f", &radius)
	area = pi * radius * radius
	fmt.Printf("radius of %g meters; area is %g square meters\n", radius, area)
}