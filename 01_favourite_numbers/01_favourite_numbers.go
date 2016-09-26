// Task 1.
// Write a function 'add' that sums two integers.
// Print a sum of two randomly chosen integers from range 0-100.

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
