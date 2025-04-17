package main

import "fmt"

/*
func init() {
	fmt.Println("Init!")
}
*/

/*
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
*/

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int, tax float64) {
	tax = 1.1
	result = int(float64(price*item) * tax)
	return result, tax
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main1() {
	// foo(10, 20)
	// foo(10, 20, 30)
	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)


	/*
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	*/

	/*
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	*/

	/*
	r,t := cal(100, 2)
	fmt.Println(r, t)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
	
	func(x int) {
		fmt.Println("inner func", x)
	}(2)
	*/

	/*
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
	*/

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

	// fmt.Println(big - 1)

	/*
	var (
		u8  uint8     = 255
		i8  int8 	  = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v\n", u8, u8)

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 + / 3 =", 10.0/3)
	fmt.Println("10 + / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)

	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	fmt.Println(1 << 0)
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)

	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println("Hello World"[0])
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	fmt.Println(s)
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	s = strings.Replace(s, "l", "L", 2)
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Hello\nWorld")
	
	fmt.Println(`Hello
	World`)

	fmt.Println("\"Hello World\"")
	fmt.Println(`"Hello World"`)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(!true)
	fmt.Println(!false)

	var y int = 1
	yy := float64(y)
	fmt.Printf("%T %v %f\n", yy, yy, yy)

	var z float64 = 1.2
	zz := int(z)
	fmt.Printf("%T %v %d\n", zz, zz, zz)

	var str string = "14"
	i, _ := strconv.Atoi(str)
	fmt.Printf("%T %v", i ,i)
	*/

	/*
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	b := [2]int{300, 400}
	b[1] = 500
	fmt.Println(b)
	fmt.Println(b[0])

	n := []int{1, 2, 3, 4, 5}
	n[1] = 100
	fmt.Println(n)
	fmt.Println(n[4])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n = append(n, 200)
	fmt.Println(n)
	n = append(n, 200, 300, 400)
	fmt.Println(n)

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
	fmt.Println(board[0])
	fmt.Println(board[0][2])

	m := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m) ,m)
	m = append(m, 1, 2, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m) ,m)

	c := make([]int, 4)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	d := []byte{72, 73}
	fmt.Println(d)
	fmt.Println((string(d)))

	e := []byte("HI")
	fmt.Println(e)
	fmt.Println((string(e)))

	ma := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(ma)
	fmt.Println(ma["apple"])
	ma["banana"] = 300
	fmt.Println(ma)
	ma["orange"] = 500
	fmt.Println(ma)
	fmt.Println(ma["mikan"])
	v, ok := ma["apple"]
	fmt.Println(v, ok)
	v2, ok2 := ma["mikan"]
	fmt.Println(v2, ok2)
	*/
}