package main

import(
  "fmt"
  "math"
)

func pow(x,n,lim float64) float64 {
  if v := math.Pow(x,n); v<lim {
    return v
  }
  return lim
}

func main() {
  fmt.Println(
    pow(3,2,10), pow(3,3,20))
}
//pow는 3,2면 3을 두 번 제곱한 값을 뜻하고
//3,3이면 3을 세 번 제곱함!
//sprt는 제곱근을 뜻하는 함수이고
//i는 허수를 뜻함 우와...엄청 나네...