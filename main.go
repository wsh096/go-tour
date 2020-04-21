package main

import"fmt"

func main(){
  s:= 1
  for s<1000 {
    s += s
  }
  fmt.Println(s)
  
}

//Go에서 "While" 사용하기
//이전의 예제에서 처럼 조건문만 표시하면 //C언어에서 while 을 사용하듯 for 를 사용할 수 있습니다.

//변수를 선언하는 const방법
//상수는 const 키워드와 함께 변수처럼 선언합니다.
//상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있습니다.