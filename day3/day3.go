package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/ahadam/aoc2018/day3/input.txt")
	var elves [1000][1000]int

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count = 0
	for scanner.Scan() {

		re := regexp.MustCompile("(\\d+)x(\\d+)")
		var l, _ = strconv.Atoi(re.FindStringSubmatch(scanner.Text())[1])
		var h, _ = strconv.Atoi(re.FindStringSubmatch(scanner.Text())[2])

		re2 := regexp.MustCompile("(\\d+),(\\d+)")
		var x, _ = strconv.Atoi(re2.FindStringSubmatch(scanner.Text())[1])
		var y, _ = strconv.Atoi(re2.FindStringSubmatch(scanner.Text())[2])

		for i := x; i < (l + x); i++ {
			for j := y; j < (h + y); j++ {
				if elves[i][j] == 2 {
					continue
				} else if elves[i][j] == 1 {
					count++
					elves[i][j] = 2
				} else {
					elves[i][j] = 1
				}
			}
		}
	}
	fmt.Println(count)
}
