package main

import "fmt"

func gcd(a, b int) int {
  if b == 0 {
    return a
  } else {
    return gcd(b, a % b)
  }
}

func main() {
  a := 10
  b := 8
  
  fmt.Printf("Greatest common divisor for %d and %d is: %d", a, b, gcd(a, b))
}