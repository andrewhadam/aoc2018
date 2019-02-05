package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("/Users/ahadam/aoc2018/day4/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile("Guard #(\\d+)begins shift")

	}
}
