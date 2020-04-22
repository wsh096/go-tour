package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  m["Answer"] =48//수정하기!
  fmt.Println("The value:", m["Answer"])

  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"]) //요소 지우기

  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok) //존재여부



  fmt.Println(m)
  }

//맵 m 의 요소를 삽입하거나 수정하기:

//m[key] = elem
//요소 값 가져오기:

//elem = m[key]
//요소 지우기:

//delete(m, key)
//키의 존재 여부 확인하기:

//elem, ok = m[key]
//위의 ok 의 값은 m 에 key 가 존재한다면 true 존재하지 않으면 false , elem 은 타입에 따라 0(zero value) 가 됩니다.