package main

import "fmt"
   //int=1, 2, 3 이렇게 입력해도 되고 안 입력해도 됨! 다 int가 아닐 수도 있으니까 안 입력하면 메모리 많이 사용

//x=1여기선 변경 x
func main() {
  x, y,z := 1, 2, 3.14
  var c, python, java = true, false, "No!"
  //x =1 여기서 각 변수의 값을 바꿔 줄 수 있다
  fmt.Println(x, y,z,c, python, java)
}

//int defalut 0
//bool(true/false) defalut 0

///이와 같이 바로 옆에서도 가능하며 bool의 경우 특정 t/f string은 ""사용!

변수의 짧은 선언인 :=는 함수 밖에서는 불가능!
즉 함수 내에서 사용하는 특정 용법이다!
하지만 var x... type = ...는 모두 사용 가능!