package main

import "fmt"

func Sqrt(x float64) float64{
  z:=float64(1)//z:=1.0
  z = z- (z*z-x)/(2*z)//z = z- (z*z-x)/(2*z)(뉴턴의 제곱근 근사값 찾는법)
  return z
}

func main(){
  fmt.Println(Sqrt(2))
}