package main

import "fmt"

var pow = [] int{1,2,4,8,16,32,64,128}

func main(){
for i, v:= range pow {
  fmt.Printf("2**%d = %d\n", i, v)
}

}

//슬라이스에서 range를 사용해 슬라이스 나 맵을 왔다 갔다 할 수 있음 ragne '앞', '뒤'
