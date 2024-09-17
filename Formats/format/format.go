package main

import(
	"fmt"
)

func main(){
	x := 34567.89

	fmt.Println("Here is Println with defaults: x = " , x)
	
	fmt.Printf("Here is Printf with f format: x = %f \n", x)
	fmt.Printf("Here is Printf with e format: x = %e \n", x)
	fmt.Printf("Here is Printf with 10g format: x = %10g \n", x)

	x = 1000 * x
	
	fmt.Printf("Here is Printf with g formats: x = %g \n", x)
	fmt.Printf("Here is Printf with 10.2g formats: x = %10.2g \n", x)
}