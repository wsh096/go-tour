package main

import "fmt"
import "math"
//이렇게 따로 쓸 수도 있고 import ("fmt" "math")

func main() {
  fmt.Printf("Now you have %g problems.",math.Nextafter(2,3))
}