// Task 2.
// Ask the user of their name and age
// and print his answers in one line.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Whats your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("How old are you? ")
	age := readInt()
	fmt.Println("Hello, " + name + ". You are " + strconv.Itoa(age) + " years old.")
}
