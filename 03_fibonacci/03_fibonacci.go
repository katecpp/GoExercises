// Task 3.
// Ask the user how many fibonacci numbers he wants to print.
// Print fibonacci sequence of that length.

package main

import "fmt"

func readInt() int {
	var i int
	for _, err := fmt.Scanln(&i); err != nil; _, err = fmt.Scanln(&i) {
		// If int read fails, read as string and forget
		var discard string
		fmt.Scanln(&discard)
		fmt.Println("This is not a valid number! Try again...")
	}
	return i
}

func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Print("How many fibonacci numbers you want to print? ")
	length := readInt()
	for i := 0; i < length; i++ {
		fmt.Print(fib(uint(i)), " ")
	}
}
