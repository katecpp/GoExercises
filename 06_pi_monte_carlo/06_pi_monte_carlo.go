// Task 6.
// Implement the estimation of PI value with the use of Monte-Carlo method.
// The number of samples used in the calculation should be passed as command line argument.

package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
	"strconv"
)

func belongsToCircle(x, y float32) bool {
	return x*x + y*y < 1
}

func estimatePi(samplesNumber int) float32 {
	rand.Seed(int64(time.Now().Second()))
	in := 0
	for i := 0; i < samplesNumber; i++ {
		randX := rand.Float32()
		randY := rand.Float32()
		if belongsToCircle(randX, randY) {
			in++
		}
	}

	var pi float32 = 4*float32(in)/float32(samplesNumber)
	return pi
}

func main() {
	if len(os.Args) == 2 {
  total, err := strconv.Atoi(os.Args[1])
  if err == nil {
		if total <= 0 {
			fmt.Println("Samples count must be greater than 0.")
			return
		}
    fmt.Println("Estimate Pi with ", total, " samples.\n")
    fmt.Println("Result: ", estimatePi(total))
    return
    }
  }
  fmt.Println("Invalid input parameters.")
  fmt.Println("Usage: 06_pi_monte_carlo.exe <samples count>")
  return
}
