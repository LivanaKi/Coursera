package main

import(
	"fmt"
)

func main(){
	var miles, yard int32 = 26, 385
	var kilometers float32

	kilometers = 1.60934 * (float32(miles) + float32(yard) / 1760.0)

	fmt.Printf("\n A marathon is %g kilometers.\n\n", kilometers)
}