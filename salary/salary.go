/* GO program that examines employee’s salaries stored in a map.The map has employee name as the key and salary as an int value. 
You wish to find and print the employee name with the smallest salary, and the employee name with the largest salary 
as well as the company’s average salary.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileWithInf = "salary.txt"

func averageSlice (sliceSalary []int64) float64{ //average salary
	var sum int64
	for _, val := range sliceSalary{
		sum += val
	}

	return float64(sum)/float64(len(sliceSalary))
}

func minSlice(sliceSalary []int64) int64 { //min salary
	var min int64
	min = sliceSalary[0]
	for _, v := range sliceSalary {
		if v <= min {
			min = v
		}
	}
	return min
}

func maxSlice(sliceSalary []int64) int64 { // max salary
	var max int64
	for _, v := range sliceSalary {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var slice []string
	var salary map[string]int64
	var name string

	salary = make(map[string]int64)

	file, err := os.Open(fileWithInf)  // open file with data
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // after implementation close file

	scanner := bufio.NewScanner(file) //scan file 
	for scanner.Scan() {   //scan the lines one by one and added in slice
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sliceSalary []int64

	for i := 0; i < len(slice); i++ {
		currInf := strings.Split(slice[i], " ")  //break the line into parts
		
		name = strings.Join(currInf[:2], "_")  // unit first name with last, used char "_"

		curSalary, _ := strconv.Atoi(currInf[len(currInf)-1]) // change type integer

		sliceSalary = append(sliceSalary, int64(curSalary))  // added in slice for look for max, min

		salary[name] = int64(curSalary) // added data in map, the name is key, the value is salary 
	}

	for key, val := range salary {
		if val == minSlice(sliceSalary) {
			fmt.Printf("Minimal salary: %s ($%d)\n", key, val)  // look for min
		}
		if val == maxSlice(sliceSalary) {
			fmt.Printf("Maximum salary: %s ($%d)\n", key, val)  // look for max
		}
	}
	fmt.Printf("The company’s average salary: $%f", averageSlice(sliceSalary))
}
