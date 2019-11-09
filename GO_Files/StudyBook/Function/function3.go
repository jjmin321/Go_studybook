//11-3

package main

import "fmt"

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
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
	//다중 값 반환

	//ex1
	a, b := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)
	fmt.Println(a, b, c, d)

	//ex2
	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, y3, _, y5 := arrayMultiply(1, 2, 3, 4, 5)
	fmt.Println(x1, x2, x3, x4, x5)
	fmt.Println(y1, y3, y5)
}
