package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var slice []name
	var nameFile string

	fmt.Println("Enter name file")
	fmt.Scanf("%s", &nameFile)

	file, err := os.Open(nameFile) // open file with data
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // after implementation close file

	scanner := bufio.NewScanner(file) //scan file
	for scanner.Scan() {              //scan the lines one by one and added in slice
		nameCurr := strings.Split(scanner.Text(), " ")
		var names name
		nameF := nameCurr[0]
		nameL := nameCurr[1]

		if len(nameF) < 20 {
			names.fname = nameF
		} else {
			names.fname = nameF[:20]
		}
		
		if len(nameL) < 20 {
			names.lname = nameL
		} else {
			names.lname = nameL[:20]
		}

		slice = append(slice, names)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
