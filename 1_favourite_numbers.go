// Task 1.
// Get two random numbers and print them.
// Print the sum of these numbers with the use of function
// that adds two integers.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func main() {
	rand.Seed(int64(time.Now().Second()))
	fav, unfav := rand.Intn(100), rand.Intn(100)
	fmt.Println("My favorite number is", fav)
	fmt.Println("My least favorite number is", unfav)
	fmt.Println("Their sum equals:", add(fav, unfav))
}
