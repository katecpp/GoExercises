// Task 3.
// Ask the user how many fibonacci numbers he wants to print.
// Print fibonacci sequence of that length.

package main

import "fmt"

func fib(x uint) uint {
	if x < 2 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Print("How many fibonacci numbers you want to print? ")
	var length int
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		fmt.Print(fib(uint(i)), " ")
	}
}
