package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	var line string
	fmt.Println("Start of the program! If you want to stop the process, type X")
	for line != "X" {
		fmt.Println("Enter integer")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line = scanner.Text()

		if line == "X" {
			fmt.Println("Stop!")
			fmt.Println("Slice:", slice)
			break
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("it is not number")
			continue
		}
		slice = append(slice, num)
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		fmt.Println(slice)
	}
}