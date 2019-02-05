package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/Users/ahadam/aoc2018/day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	is := []rune(input)

	fmt.Println("Initial array", len(is))
	var notFound = true

	for notFound {
		notFound = false
		for i := 0; i < len(is)-1; i++ {
			if is[i] == is[i+1]+32 || is[i] == is[i+1]-32 {
				notFound = true
				is = append(is[:i], is[i+2:]...)
			}
		}
	}

	// Need a way to convert the rune array back into a string array
	fmt.Println("Final array", len(is))
}
