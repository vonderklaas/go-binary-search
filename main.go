package main

import (
	"fmt"
)

func binarySearch(arr [10]int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] == target {
			return middle
		}

		if arr[middle] > target {
			right = middle - 1
		}

		if arr[middle] < target {
			left = middle + 1
		}
	}

	return -1
}

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	findFive := binarySearch(arr, 5)
	findHundread := binarySearch(arr, 5)

	fmt.Println("findFive:", findFive)
	fmt.Println("findHundread:", findHundread)
}
