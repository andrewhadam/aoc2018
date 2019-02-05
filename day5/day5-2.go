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
	var notFound = true
	var temp []rune
	var minlength = 0

	fmt.Println(len(is), "Original IS")
	for i := 'A'; i < 'Z'; i++ {
		temp = temp[:0]
		fmt.Println("Length of temp : ", len(temp))
		fmt.Println("I: ", i)
		for _, slot := range is {
			//		fmt.Println("Slot: ", slot, " i as a rune: ", rune(i), " rune -32 ", rune(i)-32, " ", j)
			if slot == rune(i) || slot == rune(i)+32 {
			} else {
				temp = append(temp, slot)
				//	fmt.Println("I am the same")
			}
		}
		fmt.Println("New temp array", len(temp), i)
		for notFound {
			notFound = false
			for k := 0; k < len(temp)-1; k++ {
				if temp[k] == temp[k+1]+32 || temp[k] == temp[k+1]-32 {
					notFound = true
					temp = append(temp[:k], temp[k+2:]...)
				}
			}
		}
		if minlength == 0 {
			minlength = len(temp)
		} else if len(temp) < minlength {
			minlength = len(temp)
		}
		print("Temp array at the end: ", len(temp), " ", i, " ")
		notFound = true
	}

	// Need a way to convert the rune array back into a string array
	fmt.Println("Final array", len(is))
	fmt.Println(minlength, " min length")

}
