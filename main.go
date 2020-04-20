package main

import "fmt"

func add(x , y int) int {//int와 같은 type이 뒤에 온다!!//따로 쓸 수도 있고 같으면 뒤에 하나만 쓸 수도 있다.
  return x + y
}

func main() {
  fmt.Println(add(42,13))
}