package main;

import "fmt"

func main()  {
  boolVar := make(chan bool, 1)
  boolVar <- true;
  go func ()  {

  }()

  fmt.Print(<-boolVar)

}
