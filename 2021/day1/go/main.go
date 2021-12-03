package main

import (
	"fmt"
	"io"
	"os"
)

func readInput(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	return nums
}

func firstPuzzle(input []int) int {
	increasedCount := 0

	for i, number := range input {
		if i == 0 {
			continue
		}

		if number > input[i-1] {
			increasedCount = increasedCount + 1
		}
	}

	return increasedCount
}

func secondPuzzle(input []int) int {

	increasedCount := 0

	// create a sliding window for our data
	// a sliding window is

	for idx, number := range input {

		// we loop through all the numbers, we need to create window of 3
		// we will skip so we can create full windows
		if idx <= 2 {
			continue
		}

		windowAggLast := input[idx-1] + input[idx-2] + input[idx-3]
		windowAggCurrent := number + input[idx-1] + input[idx-2]

		if windowAggCurrent > windowAggLast {
			increasedCount = increasedCount + 1
		}

	}

	return increasedCount

}

func main() {

	nums := readInput("input.txt")

	puzzleOne := firstPuzzle(nums)
	fmt.Printf("First puzzle: %d\n", puzzleOne)

	puzzleTwo := secondPuzzle(nums)
	fmt.Printf("Second puzzle: %d\n", puzzleTwo)

	os.Exit(1)
}
