package main

import(
  "fmt"
)

func main(){
  j := 0;
  for j < 4 {
  for i:=0; i< 4; i++ {
    fmt.Printf("_%d", i);
    fmt.Print("_");
    fmt.Print("+");
  }
  j++;
  fmt.Println();
}
}
