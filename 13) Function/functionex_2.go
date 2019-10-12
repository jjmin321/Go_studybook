package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	//함수 고급
	//함수를 변수에 할당

	//ex1 (슬라이스에 할당)
	f := []func(int, int) int{multiply, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println(a)
	fmt.Println(b)

	//ex2 (변수에 할당)
	f1 := multiply
	f2 := sum
	fmt.Println(f1(0, 5), f2(0, 5))

	//ex3 (맵에 할당)
	m := map[string]func(int, int) int{
		"multiply": multiply,
		"sum":      sum,
	}

	fmt.Println(m["multiply"](10, 10))
	fmt.Println(m["sum"](10, 10))
}
