package main

import"fmt"

const Pi = 3.14

func main(){
  const W = "안녕"
  fmt.Println("Hello", W)
  fmt.Println("Happy",Pi,"Day")

  const T = true
  fmt.Println("Gp rules?", T)
}

//변수를 선언하는 const방법
//상수는 const 키워드와 함께 변수처럼 선언합니다.
//상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있습니다.