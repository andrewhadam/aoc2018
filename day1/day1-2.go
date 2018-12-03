package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var m map[int]int
	m = make(map[int]int)
	file, err := os.Open("/Users/ahadam/aoc2018/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count = 0
	var found = false
	m[count] = 1

	for !found {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			toI, _ := strconv.Atoi(scanner.Text())

			if _, ok := m[count]; ok {
				if count == 0 {
					fmt.Println("blah")
				} else {
					fmt.Println("First duplicate")
					found = true
					break
				}
			} else {
				m[count] = 1
			}
			count += toI
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(count)

	}
}
