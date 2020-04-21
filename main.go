package main

import "fmt"
        
type Vertex struct {//X,Y int}로 써도 됨 
  X int
  Y int
  }

func main() {
  v := Vertex{1,2}
  w := &v
  w.X =1e9//.을 이용해서 X의 값에 접근해서 변경할 수도 있고 그걸 확인 할 수도 있음! 신기하구만!

fmt.Println(*w)//*을 찍어서 그 주소의 값을 가져오라는 명령문을 실행시켜야 한다. 안 그러면 앞에 불필요한 &가 붙는다.
fmt.Println(v)  

}