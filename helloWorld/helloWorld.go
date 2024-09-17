package main

import (
	"fmt"
	//"strings"
)
func main() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x:=0; x<5; x++ {
	  xtemp = x2 //1 2 3 5
	  x2 = x2 + x1//1  2  3  5 8
	  x1 = xtemp //1  1 2 3
	}
	fmt.Println(x2)
  }
  