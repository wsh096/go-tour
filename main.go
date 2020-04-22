package main

import "fmt"

func main(){
var z [] int
fmt.Println(z, len(z),cap(z))
if z== nil {
  fmt.Println("nil!")
}
}

//빈 슬라이스의 기본 값은 nil인 0,0
