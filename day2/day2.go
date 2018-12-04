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
		for len(str) > 0 {
			//for ; spos, char := range str {
			fmt.Println("Removed chat", string(str[0]))
			if strings.Count(str, string(str[0])) == 2 {
				twoChar++
				fmt.Println("Add two")
			} else if strings.Count(str, string(str[0])) == 3 {
				threeChar++
				fmt.Println("Add three")
			} else {
				fmt.Println("Neither")
			}
			str = strings.Replace(str, string(str[0]), "", -1)
			fmt.Println("String post manipulation", str)
			//			str = strings.Replace(str, string(char), "", -1)
			fmt.Println("Replacement String: ", str)
			//fmt.Println(pos, char)
		}
		//		var string = "blahblah"
		//		duplicates = string.chars.group_by { |char| char }.select { |key, value| value.size > 1 }

	}
	fmt.Println(twoChar, threeChar)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
