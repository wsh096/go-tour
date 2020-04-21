package main

import(
  "fmt"
  "math"
)

func sqrt(x float64) string{
  if x<0{
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}
func main() {
  fmt.Println(sqrt(2),sqrt(-4))
}

//sprt는 제곱근을 뜻하는 함수이고
//i는 허수를 뜻함 우와...엄청 나네...