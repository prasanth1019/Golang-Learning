package main

import "fmt"

type Expr func() float64

func main() {
	x := 1.0
	// ex := Const(32)
	fmt.Println(&x)
	fmt.Println(Ref(&x))
}

func Const(x *float64) Expr {
	return func() float64 {
		return *x
	}
}

func Ref(x float64) Expr {
	return func() float64 {
		return x
	}
}
