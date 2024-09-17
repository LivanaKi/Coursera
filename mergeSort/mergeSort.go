/*
GO program that does mergesort. We use one slice that starts with the unordered data. 
The program generate appropriate int test data pseudo-randomly (10,000 elements)
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	var mergedSlice []int
	return merge(first, second, mergedSlice)
}

func merge(firstSlice []int, secondSlice []int, mergedSlice []int) []int {
	var index1, index2 int
	for index1 < len(firstSlice) && index2 < len(secondSlice) {
		if firstSlice[index1] < secondSlice[index2] {
			mergedSlice = append(mergedSlice, firstSlice[index1])
			index1++
		} else {
			mergedSlice = append(mergedSlice, secondSlice[index2])
			index2++
		}
	}
	for ; index1 < len(firstSlice); index1++ {
		mergedSlice = append(mergedSlice, firstSlice[index1])
	}
	for ; index2 < len(secondSlice); index2++ {
		mergedSlice = append(mergedSlice, secondSlice[index2])
	}
	return mergedSlice
}

func writeFile(fileName string, dataSlice []int){
	file, err := os.Create(fileName)
	checkFile(err)
	defer file.Close()

	for _, v := range dataSlice{
		_, err := fmt.Fprintf(file, "%d\n", v)

		checkFile(err)
	}
}
func checkFile (e error){
	if e != nil{
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var size = 10000

	dataSlice := make([]int, size)
	for i, _ := range dataSlice {
		dataSlice[i] = rand.Intn(10000)
	}
	fmt.Println("Start program\nTime:", time.Now())
	writeFile("sortedData.num", mergeSort(dataSlice))
	fmt.Println("Stop program\nTime:", time.Now())
	fmt.Println("Length slice:", len(dataSlice))
	

}
