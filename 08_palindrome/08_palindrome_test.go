package main

import "testing"

func TestPalindrome(t *testing.T){
  if isPalindrome("") != true {
    t.Error("Test failed")
  }
  if isPalindrome("a") != true {
    t.Error("Test failed")
  }
  if isPalindrome("aa") != true {
    t.Error("Test failed")
  }
  if isPalindrome("Aaa") != true {
    t.Error("Test failed")
  }
  if isPalindrome("abba") != true {
    t.Error("Test failed")
  }
  if isPalindrome("aabb") != false {
    t.Error("Test failed")
  }
  if isPalindrome("longlonglongstring") != false {
    t.Error("Test failed")
  }
}
