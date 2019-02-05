package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type rect struct {
	x       int
	y       int
	l       int
	h       int
	overlap bool
	index   int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("/Users/ahadam/aoc2018/day3/input.txt")
	var elves [1000][1000]int
	var moreelves []rect

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

		re3 := regexp.MustCompile("#(\\d+)")
		var id, _ = strconv.Atoi(re3.FindStringSubmatch(scanner.Text())[1])

		newObj := rect{x, y, l, h, false, id}

		// initiated as 0, 1 is I'm only there, 2 is I overlap something
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

		moreelves = append(moreelves, newObj)
	}

	for i, elf := range moreelves {
		for j, repelf := range moreelves {
			if i == j {
				continue
			}
			if repelf.overlap == true {
				continue
			}
			if (min(elf.x+elf.l, repelf.x+repelf.l)-max(elf.x, repelf.x) > 0) && (min(elf.y+elf.h, repelf.y+repelf.h)-max(elf.y, repelf.y) > 0) {
				moreelves[j].overlap = true
				moreelves[i].overlap = true
			}
		}
	}
	var counter = 0
	for _, things := range moreelves {
		if things.overlap == false {
			fmt.Println("found: ", things)
			counter++
		}
	}
	fmt.Println(counter)
}
