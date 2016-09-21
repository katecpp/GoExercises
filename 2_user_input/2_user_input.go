// Task 2.
// Ask the user of their name and age
// and print his answers in one line.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Whats your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("How old are you? ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)

	fmt.Println("Hello, " + name + ". You are " + age + " years old.")
}
