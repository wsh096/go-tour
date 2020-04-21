package main

import"fmt"

func main(){
  s:= 0
  for i:=0; i<10; i++{
    s += i
  }
  fmt.Println(s)
  
}

//변수를 선언하는 const방법
//상수는 const 키워드와 함께 변수처럼 선언합니다.
//상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있습니다.