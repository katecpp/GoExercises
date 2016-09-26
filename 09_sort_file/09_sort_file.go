// Task 9.
// Read data from file containing some integers, each in new line.
// Sort the data in ascending order and write sorted numbers into new file.
// Each number should be in newline, just like in input file.

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readIntsFromFile(fileName string) []int {
  var numbers []int
  f, err := os.Open(fileName)
  defer f.Close()
  check(err)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
      lineStr := scanner.Text()
      num, _ := strconv.Atoi(lineStr)
      numbers = append(numbers, num)
  }
  return numbers
}

func writeIntsToFile(numbers []int, fileName string){
  f, err := os.Create(fileName)
  defer f.Close()
  check(err)
  writer := bufio.NewWriter(f)
  for i := 0; i < len(numbers); i++ {
    fmt.Fprint(writer, strconv.Itoa(numbers[i]) + "\r\n")
  }
  writer.Flush()
}

func main() {
  numbers := readIntsFromFile("input_data_numbers.txt")
  sort.Ints(numbers)
  writeIntsToFile(numbers, "output_data_numbers.txt")
  fmt.Println("Done.")
}
