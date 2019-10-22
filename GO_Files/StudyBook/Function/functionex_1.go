//함수 심화
//함수 내에 많은 변수를 넣고 많은 변수를 반환할 때 사용하는 문법 //나의 주관으로는 거의 사용하지 않음.
package main

import "fmt"

func multiply(n ...int) int {
	total := 1
	for i, value := range n {
		fmt.Println("multiply = ", i, value)
		total *= value
	}

	return total
}

func sum(n ...int) int {
	total := 0
	for i, value := range n {
		fmt.Println("sum = ", i, value)
		total += value
	}
	return total
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
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
	//가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	//ex1
	x := multiply(5, 6, 7, 8, 9, 10) //x[0] = 5 , x[1] = 6, x[2] = 7 ..... 순으로 전달됨.
	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("multiply = ", x)
	fmt.Println("sum = ", y)

	//ex2
	prtWord("a", "b", "ab", "banana", "golang", "hello world!")

	//ex3
	a := []int{1, 2, 3, 4, 5}

	m := multiply(a...)
	n := sum(a...)
	fmt.Println("m = ", m)
	fmt.Println("n = ", n)

}
