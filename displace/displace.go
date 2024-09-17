package main

import(
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, velocity, displacement float64) func (float64) float64{
	
	fn := func (time float64) float64{
		return ((0.5) * acceleration * math.Pow(time, 2) + velocity * time + displacement)
	}
	
	return fn
}

func main(){
	var acceleration, velocity, displacement, time float64
	fmt.Print("Enter acceleration: ")
	fmt.Scanf("%f\n", &acceleration)
	fmt.Print("Enter velocity: ")
	fmt.Scanf("%f\n", &velocity)
	fmt.Print("Enter displacement: ")
	fmt.Scanf("%f\n", &displacement)

	fn := GenDisplaceFn(acceleration, velocity, displacement)

	fmt.Print("Enter time: ")
	fmt.Scanf("%f\n", &time)
	fmt.Println("Displacement:", fn(time))

	fmt.Print("Enter another time: ")
	fmt.Scanf("%f\n", &time)
	fmt.Println("Displacement:", fn(time))
}