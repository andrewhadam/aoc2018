package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CompareInsensitive(a string, b string) bool {
	// loop over string a and convert every rune to lowercase
	var sa = strings.Split(a, "")
	var sb = strings.Split(b, "")

	var diffcount = 0

	for i := 0; i < len(sa); i++ {
		if sa[i] != sb[i] {
			diffcount++
		}
		if diffcount > 1 {
			return false
		}
	}

	if diffcount == 1 {
		fmt.Println(sa)
		fmt.Println(sb)
		fmt.Println(diffcount)
		fmt.Println("Break-------")
		return true
	} else {
		return false
	}
}

func main() {

	file, err := os.Open("/Users/ahadam/aoc2018/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputs []string

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, str := range inputs {
		for j := 0; j < len(inputs); j++ {
			CompareInsensitive(str, inputs[j])
		}
	}
}
