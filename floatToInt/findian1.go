package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)
/*
func main() {

  fmt.Println("Enter a string")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  line := scanner.Text()

  line = strings.ToLower(line)

  if strings.HasPrefix(line, "i") && strings.Contains(line, "a") && strings.HasSuffix(line, "n") {
    fmt.Println("Found!")
  } else {
  	fmt.Println("Not Found!")
  }

}