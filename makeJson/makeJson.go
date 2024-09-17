package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string

	mapping := make(map[string]string)

	fmt.Println("Enter name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	mapping["name"] = name
	
	fmt.Println("Enter address")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	address = scanner2.Text()
	mapping["address"] = address

	js, err := json.Marshal(mapping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}
