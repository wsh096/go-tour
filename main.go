package main

import "fmt"

func main(){
  a:= make([]int,5)//5개의 빈 배열 만듬(길이도 5 용량도 5)
  printSlice("a",a)
  b:= make([]int,0,5)
  printSlice("b",b)//길이는 0, 용량은 5인 배열을 만듬
  c:=b[:2]
  printSlice("c",c)//?길이는 2인데//용량은 5
  d:=b[2:5] //길이3,용량3
  printSlice("d",d)
}
func printSlice(s string, x[ ]int) {
  fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}
  
//슬라이스 자르기