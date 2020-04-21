package main

import(
  "fmt"
  "math"
)

func pow(x,n,lim float64) float64 {
  if v := math.Pow(x,n); v<lim {
    return v
  }else{//위의 조을 else도 그대로 받음!
  //즉, v := math.Pow(x,n); v<lim이외의 조건 v>=lim이 자동으로 입력
    fmt.Printf("%g >= %g\n",v,lim)
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

//if 에서 짧은 명령문을 통해 선언된 변수는 else 블럭 안에서도 사용할 수 있습니다.

//%g를 하면 플롯을 적당한 정수 값으로 표현해줌