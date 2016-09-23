// Task 10.
// Implement Caesar cipher, which takes two command line arguments:
// Input string to decode and the offset value for Caesar algorithm.

package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
)

func invalidArgsMsg() {
  fmt.Println("Invalid input parameters.")
  fmt.Println("Usage: 10_ceasar.exe <string to decode> <offset>")
  return
}

func cipher(s string, offset int) string {
  const alphabetLen int = 26
  s = strings.ToLower(s)
  offset = offset % alphabetLen
  result := ""
  for i := 0; i < len(s); i++ {
      ascii := int(s[i]) + offset
      if ascii > int('z') {
        ascii -= alphabetLen
      } else if ascii < int('a') {
        ascii += alphabetLen
      }
      result += string(ascii)
  }
  return result
}

func main() {
  if len(os.Args) == 3 {
  stringToDecode := strings.ToLower(os.Args[1])
  offset, err := strconv.Atoi(os.Args[2])
  if err == nil {
    fmt.Printf("Encode \"%s\" with offset %d \r\n", stringToDecode, offset)
    fmt.Println("Result: ", cipher(stringToDecode, offset))
    return
    }
  }
  fmt.Println("Invalid input parameters.")
  fmt.Println("Usage: 10_ceasar.exe <string to decode> <offset>")
  return
}
