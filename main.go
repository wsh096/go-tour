package main

import( "fmt"
        "runtime"
)


func main(){
  
    fmt.Println("Go runs on ")
    switch os := runtime.GOOS; os {
      case "darwin" :
      fmt.Println("OS X.")
      case "linux":
      fmt.Println("Linux")
      default:
      fmt.Printf("%s",os)
    }
  }

//그리고 함수는 클로져(full closures) 입니다.

//코드에서 adder 함수는 클로져(closure)를 반환합니다.

//각각의 클로져는 자신만의 sum 변수를 가집니다.