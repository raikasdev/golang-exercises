package main

import (
	"fmt"
	"math"
	"strconv"
)

// Golang Binary Search
// (C) 2022 Roni Äikäs under Expat/MIT license
var numbers = [25]int{1, 3, 8, 9, 11, 13, 21, 22, 24, 26, 29, 32, 36, 45, 58, 60, 61, 64, 69, 71, 78, 80, 90, 96, 99}

func main() {
	// Println function is used to
	// display output in the next line
	fmt.Println("Number to find from the array: ")
	var input string
	fmt.Scanln(&input)
	number, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	find(int(number), 0, int64(len(numbers)-1))
	fmt.Println("Hello, World!")
}

func find(target int, start int64, end int64) {
	if start > end {
		fmt.Println("Target not in array")
		return
	}
	var middleIndex int64 = int64(math.Floor(float64((start + end) / 2)))
	if numbers[middleIndex] == target {
		fmt.Println("Target was in index", middleIndex)
		return
	}
	if numbers[middleIndex] > target {
		fmt.Println("Target is in range ", start, "-", middleIndex-1)
		find(target, start, middleIndex-1)
	} else {
		fmt.Println("Target is in range ", middleIndex+1, "-", end)
		find(target, middleIndex+1, end)
	}
}
