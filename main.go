package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string] Vertex{
  "Bell Labs":{
    40.68433, -74.39967,
    },
  "Google": {
    37.42202,-122.08408,
    },
  }

func main() {
  
  fmt.Println(m)
  }
//맵의 리터럴
//맵의 문자열

//가장 상위의 타입이 타입명이라면 문자열의 타입명 생략이 가능
//Vertex생략가능