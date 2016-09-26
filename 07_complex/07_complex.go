// Task 7.
// Create a structure that represents complex number.
// Implement mathematical operations for complex as a methods:
// Add, subtract, multiply.

package main

import (
  "fmt"
)

type Complex struct {
  r float64
  i float64
}

func (x *Complex) print() {
  fmt.Println(x.r, "+", x.i, "i")
}

func (x *Complex) add(y Complex) Complex {
  return Complex{x.r + y.r, x.i + y.i}
}

func (x *Complex) sub(y Complex) Complex {
  return Complex{x.r - y.r, x.i - y.i}
}

func (x *Complex) mult(y Complex) Complex {
  return Complex{x.r*y.r - x.i*y.i, x.r*y.i + x.i*y.r}
}

func main() {
  a := Complex{1, 1}
  b := a.add(Complex{1, -1})
  c := a.mult(b)
  d := c.sub(b)

  b.print()
  c.print()
  d.print()
}
