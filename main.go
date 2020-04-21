package main

import "fmt"

type Vertex struct {
  X,Y int
}

func main(){
  v:= new(Vertex)
  fmt.Println(*v)//포인터여서 *안 붙이면 &가 붙어서 나옴!!!
  v.X, v.Y = 11, 9
  fmt.Println(*v) 
}



//new(T) 는 모든 필드가 0(zero value) 이 할당된 T 타입의 !포인터를 반환합니다!