// Task 6.
// Estimate PI value with Monte-Carlo method.

package main

import (
	"math/rand"
	"fmt"
	"time"
)

func belongsToCircle(x, y float32) bool {
	return x*x + y*y < 1
}

func main() {
	rand.Seed(int64(time.Now().Second()))

	in := 0
	const total = 1000

	for i := 0; i < total; i++ {
		randX := rand.Float32()
		randY := rand.Float32()
		if belongsToCircle(randX, randY) {
			in++
		}
	}

	var pi float32 = 4*float32(in)/float32(total)
	fmt.Println("Estimated PI value: ", pi)
}
