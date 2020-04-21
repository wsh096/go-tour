package main

import "fmt"
        
type Vertex struct {//X,Y int}로 써도 됨 
  X int
  Y int
  }

func main() {
  fmt.Println(Vertex{1,2})
}