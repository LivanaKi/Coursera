package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func sortSlice(wg *sync.WaitGroup, slice []int, c chan []int) {
	defer wg.Done()
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	c <- slice
}

func partsSlice(numParts int, sourceSlice []int) [][]int {
	slices := make([][]int, numParts)
	arrayLen := len(sourceSlice)
	partsLen := arrayLen / numParts
	
	for i := 0; i < numParts; i++ {
		start := i * partsLen
		end := start + partsLen
		if i == numParts-1 {
			end = arrayLen
		}
		slices[i] = sourceSlice[start:end]
	}
	return slices
}

func atoi(line string) []int{
	var sourceSlice []int
	lineSeparated := strings.Split(line, " ")
	for _, val := range lineSeparated {
		counter, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("it is not number")
			break
		}
		sourceSlice = append(sourceSlice, counter)
	}
	return sourceSlice
}

func unitedElements (num []int, sliceUnited []int) []int{
	for i := 0; i < len(num); i++ {
		sliceUnited = append(sliceUnited, num[i])
	}
	return sliceUnited
}

func main() {
	var line string
	var sliceUnited []int
	c := make(chan []int)
	var num []int

	numParts := 4
	
	fmt.Println("Enter a series of integers separated by spaces: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()

	sourceSlice := atoi(line)
	slices := partsSlice(numParts, sourceSlice)

	for i := 0; i < numParts; i++ {
		wg.Add(1)
		go sortSlice(&wg, slices[i], c)
		num = <-c // new code here
		fmt.Printf("slice %d\n", num)
		sliceUnited = unitedElements(num, sliceUnited)
	}
	
	wg.Wait()
	close(c)
	
	sort.Slice(sliceUnited, func(i, j int) bool {
		return sliceUnited[i] < sliceUnited[j]
	})
	
	fmt.Println("Finished:", sliceUnited)
}
