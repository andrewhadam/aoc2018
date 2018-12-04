package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("/Users/ahadam/aoc2018/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var twoChar = 0
	var threeChar = 0

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		var str = scanner.Text()
		fmt.Println("OG String: ", str)
		var twoTrue = false
		var threeTrue = false
		for len(str) > 0 {
			fmt.Println("Removed chat", string(str[0]))
			if strings.Count(str, string(str[0])) == 2 && twoTrue == false {
				twoChar++
				twoTrue = true
				fmt.Println("Add two")
			} else if strings.Count(str, string(str[0])) == 3 && threeTrue == false {
				threeChar++
				threeTrue = true
				fmt.Println("Add three")
			} else {
				fmt.Println("Neither")
			}
			str = strings.Replace(str, string(str[0]), "", -1)
			fmt.Println("String post manipulation", str)
			fmt.Println("Replacement String: ", str)
		}
	}
	fmt.Println(twoChar, threeChar)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
