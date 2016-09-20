// Task 5.
// Generate an array of 20 random values from range 0-100. Print it.
// Then search for the minimum value in this array and print it.
package main

import (
	"math/rand"
	"time"
	"fmt"
)

func getRandomArray() []int {
	rand.Seed(int64(time.Now().Second()))
	const size = 20
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func findMinimum(arr []int) int {
	if len(arr) == 0 {
		fmt.Println("Array is empty.")
		return -1
	}

	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}

func main() {
	randomArray := getRandomArray()
	fmt.Println(randomArray)
	minValue := findMinimum(randomArray)
	fmt.Println("Minumim: ", minValue)
}
