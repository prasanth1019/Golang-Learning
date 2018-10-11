package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
      )

type objStruct struct {
  gender string   `json:"results>gender"`
  email string    `json:"results>email"`
}

type manStruct struct {
  Results []struct {
    Gender string `json:"gender"`
    Email string `json:"email"`
  } `json:"results"`
}

func main() {
  var ms manStruct
  const url = "https://randomuser.me/api?results=1"
  res, _ :=  http.Get(url)
  bytes, _ :=  ioutil.ReadAll(res.Body)
  // fmt.Println(string(bytes))
  json.Unmarshal(bytes, &ms)
  fmt.Println("your data")
  fmt.Printf("%+v", ms)
  // for k, v := range ms {
  //   fmt.Println(v)
  // }
}

// type response2 struct {
//     Page   int      `json:"page"`
//     Fruits []string `json:"fruits"`
// }
//
// func main()  {
//   str := `{"page": 1, "fruits": ["apple", "peach"]}`
//   res := response2{}
//   json.Unmarshal([]byte(str), &res)
//   fmt.Printf("%+v", res)
// }
