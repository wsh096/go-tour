package main

import "fmt"

func main(){
  pow := make([]int, 10)
for i := range pow {
  pow[i] = 1<< uint(i)
}
for i,_ := range pow { //_가 있어야 그 값의 빈칸으로 인식한다!!
  fmt.Printf("%d\n",i)
}
}

//슬라이스에서 range를 사용해 슬라이스 나 맵을 왔다 갔다 할 수 있음 ragne '앞', '뒤'
