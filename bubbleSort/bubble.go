package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(d []int) {
	for i := 0; i < len(d)-1; i++ {
		Swap(i, d)
	}
}

func Swap(id int, d []int) {
	for j := 0; j < len(d)-id-1; j++ {
		if d[j] > d[j+1] {
			d[j], d[j+1] = d[j+1], d[j]
		}
	}
}

func stringToInt(lineSlice []string, slice []int) []int{
	if len(lineSlice) > 10 {
		lineSlice = lineSlice[:10]
		fmt.Println("You enter more number!")
	}

	if len(lineSlice) < 10{
		fmt.Println("You enter less number!")
	}

	for i := 0; i < len(lineSlice); i++ {
		num, err := strconv.Atoi(lineSlice[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		slice = append(slice, num)
	}
	return slice
}

func main() {
	var slice []int

	fmt.Println("Enter 10 numbers consecutively")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	lineSlice := strings.Split(line, " ")
	slice = stringToInt(lineSlice, slice)

	fmt.Println("Slice:", slice)
	bubbleSort(slice)
	fmt.Println("Bubble sort slice:", slice)
}
