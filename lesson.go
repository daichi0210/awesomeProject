package main

import (
	"fmt"
)

/*
func init() {
	fmt.Println("Init!")
}
*/

func bazz() {
	fmt.Println("Bazz")
}

var (
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t, f bool = true, false
)

func foo() {
	xi := 2
	xf64 := 1.3
	xs := "test test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Println(i, f64, s, t, f)
}

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

const big = 9223372036854775807 + 1

func main() {
	//bazz()

	/*
	fmt.Println("Hello world!", "golang")
	fmt.Println(time.Now())
	fmt.Println(user.Current())
	*/
	
	/*
	var i int = 1
	var t,f bool = true,false
	fmt.Println(i)
	fmt.Println(t)
	fmt.Println(f)
	*/

	/*
	var (
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		t, f bool = true, false
	)
	fmt.Println(i, f64, s, t, f)
	*/

	/*
	var (
		i int
		f64 float64
		s string
		t, f bool
	)
	fmt.Println(i, f64, s, t, f)
	*/

	/*
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", xf64)
	*/

	/*
	fmt.Println(i, f64, s, t, f)
	foo()
	*/

	// fmt.Println(Pi, Username, Password)

	fmt.Println(big - 1)
}