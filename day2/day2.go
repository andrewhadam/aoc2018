package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/Users/ahadam/aoc2018/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
