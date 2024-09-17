package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("size of slice is")
	size := 0
	fmt.Scanf("%d", &size)
	data := make([]float64, size)
	for i := 0; i < size; i++{
		data[i] = rand.Float64()*1000.0
	}
	startTime := time.Now()
	fmt.Println("min of", size, "rand.Float64() is", minSlice(data))
	fmt.Println("max of", size, "rand.Float64() is", maxSlice(data))
	duration := time.Since(startTime)
	fmt.Println("running time of", duration)

}

func minSlice(d []float64) float64{
	var min float64 = 1000.0
	for _, v := range d{
		if v < min{
			min = v
		}else{
			v = min
		}
	}
	return min
}

func maxSlice(d []float64) float64{
	var max float64 = 0.0
	for _, v := range d{
		if v > max{
			max = v
		}
	}
	return max
}