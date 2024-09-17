package main

import(
	"fmt"
	"math"
)

func main(){
	var area, radius, volume, circumference float64

	fmt.Println("Input radius")
	fmt.Scanf("%f", &radius)

	area = math.Pi * radius * radius
	volume = (4.0 / 3.0) * math.Pi * math.Pow(radius, 3.0)
	circumference = 2 * math.Pi * radius
	
	fmt.Printf("radius of %g meters; area is %g square meters; volume is %g; circumference is %g\n", radius, area, volume, circumference)


}