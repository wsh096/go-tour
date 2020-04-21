package main

import "fmt"
        
type Vertex struct {//X,Y int}로 써도 됨 
  X int
  Y int
  }

func main() {
  v := Vertex{1,2}
  v.X =4//.을 이용해서 X의 값에 접근해서 변경할 수도 있고 그걸 확인 할 수도 있음! 신기하구만!
  fmt.Println(v.X)
}