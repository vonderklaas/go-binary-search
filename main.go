package main

import "fmt"

func main() {
	sls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	linearSearch(sls, 10)
	binarySearch(sls, 10)
}

func linearSearch(arr []int, target int) bool {
	for i, value := range arr {
		if value == target {
			fmt.Printf("linearSearch runs: %d | %s\n", i+1, "true")
			return true
		}
	}
	fmt.Println("linearSearch: not found")
	return false
}

func binarySearch(arr []int, target int) bool {
	leftPointer := 0
	rightPointer := len(arr) - 1
	counter := 0

	for leftPointer <= rightPointer {
		counter++
		middlePointer := (rightPointer + leftPointer) / 2

		if arr[middlePointer] == target {
			fmt.Printf("binarySearch runs: %d | %s\n", counter, "true")
			return true
		}

		if arr[middlePointer] > target {
			rightPointer = middlePointer - 1
		} else {
			leftPointer = middlePointer + 1
		}
	}

	fmt.Println("binarySearch: not found")
	return false
}
