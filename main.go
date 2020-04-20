package main

import "fmt"

func swap(x , y string) (string, string){//int와 같은 type이 뒤에 온다!!//따로 쓸 수도 있고 같으면 뒤에 하나만 쓸 수도 있다.
  return y, x
}

func main() {
  a, b:= swap("Hello", "World")
  fmt.Println(a,b)
}