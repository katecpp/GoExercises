// Task 8.
// Create a function that checks if string is a palindrome.
// Function should ignore letter case.
// Create tests for this function.

package main

import (
  "fmt"
  "strings"
)

func isPalindrome(s string) bool {
  str := strings.ToLower(s)
  for beginIdx, endIdx := 0, len(s) - 1; beginIdx < endIdx; {
    if str[beginIdx] != str[endIdx] {
      return false
    }
    beginIdx++
    endIdx--
  }
  return true
}

func main() {
  a := "kobylamamalybok"
  fmt.Println(isPalindrome(a))
}
