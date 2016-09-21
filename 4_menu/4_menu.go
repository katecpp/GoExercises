// Task 4.
// Create a simple menu that continuously asks a user for choice.
// It should do following actions for given input:
// 1 - Print current date
// 2 - Print current time
// 3 - Print current weather info
// 4 - Exit

package main

import (
	"fmt"
	"time"
)

func printMenu() {
	fmt.Println("==========================")
	fmt.Println("Choose 1 to get current date")
	fmt.Println("Choose 2 to get current time")
	fmt.Println("Choose 3 to get current weather")
	fmt.Println("Choose 4 to exit")
	fmt.Println("==========================")
}

func printTime() {
	date := time.Now().Local().Format("3:04PM")
	fmt.Println("Current time is: ", date)
}

func printDate() {
	date := time.Now().Local().Format("2006-01-02")
	fmt.Println("Today is: ", date)
}

func printWeather() {
	hour := time.Now().Local().Hour()
	if hour > 8 && hour < 20 {
		fmt.Println("Sun is shining!")
	} else {
		fmt.Println("It's very dark.")
	}
}

func main() {
	fmt.Println("==========================")
	fmt.Println("Welcome to Go Menu Choice!")

	var choice int

	for choice != 4 {
		printMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			printDate()
		case 2:
			printTime()
		case 3:
			printWeather()
		case 4:
			fmt.Println("Bye!")
		default:
			fmt.Printf("Are you sure you have chosen properly?")
		}
	}
}
