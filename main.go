package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    fn0 := 0

    fn1 := 0

    return func() int{
      fn2:=fn1 +fn0
      fn1, fn0 = fn2, fn1
      if fn2 == 0 {

            fn1 = 1
      }
      return fn2
    }
}

func main(){
  f := fibonacci()
  for i := 0; i <10; i++ {
    fmt.Println(f())
  }
}
//그리고 함수는 클로져(full closures) 입니다.

//코드에서 adder 함수는 클로져(closure)를 반환합니다.

//각각의 클로져는 자신만의 sum 변수를 가집니다.