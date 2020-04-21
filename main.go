package main

import"fmt"

const (
  Big = 1<<100//2진수의 1을 100자리 오른쪽으로 이동
  Small = Big>>99//big을 99자리 땡겨옴
  //2진수상 0001이 0010이므로 2임
)//const값을 위와 같이 범위로 정의 가능

func needInt(x int) int{ 
  return x*10+1
 }
func needFloat(x float64) float64 {
  return x*0.1
}

func main(){
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}

//변수를 선언하는 const방법
//상수는 const 키워드와 함께 변수처럼 선언합니다.
//상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있습니다.