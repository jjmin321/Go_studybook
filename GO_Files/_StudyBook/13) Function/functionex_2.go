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
	//Function
	//선언 : func 키워드로 선언
	//func 함수명(매개변수) (반환타입 or 반환 값 변수명)
	//func 함수명() (반환타입 or 반환 값 변수명)
	//func 함수명(매개변수)
	//func 함수명()
	//타 언어와 달리 return value 무한적으로 가능
	//변수에 함수 넣기 -> return 사용 시 변수에 value , 미사용시 변수는 function의 키가 됨.
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
