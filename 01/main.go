package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	var numSlice []string
	var globalCount int

	// Read in file line by line
	file, err := os.Open("numlist.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		numSlice = append(numSlice, fileScanner.Text())
	}

	file.Close()

	// Iterate through list and call getNum func to determine number in line
	// Add that number to global number.
	for _, line := range numSlice {
		globalCount = globalCount + extractNum(line)
	}

	// Print global number
	fmt.Println(globalCount)
}

func extractNum(line string) int {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var foundNums []string
	var returnNum int
	var convString string

	for _, char := range line {
		if slices.Contains(numbers, string(char)) {
			foundNums = append(foundNums, string(char))
		}
	}

	convString = fmt.Sprintf("%s%s", foundNums[0], foundNums[len(foundNums)-1])

	returnNum, err := strconv.Atoi(convString)
	if err != nil {
		log.Fatal(err)
	}

	return returnNum
}
