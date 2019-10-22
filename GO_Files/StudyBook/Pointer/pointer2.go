// 출력을 해보고 식과 답을 비교해보시면 포인터의 기초에 대해 이해하기 쉽습니다.

package main

import "fmt"

func main() {
	//ex1
	i := 7
	p := &i

	fmt.Println(i, *p)
	fmt.Println(&i, p)

	*p++ //1증가

	fmt.Println(i, *p)
	fmt.Println(&i, p)

	*p = 77777 //포인터 변수 역 참조 값 변경

	fmt.Println(i, *p)
	fmt.Println(&i, p)

	i = 77

	fmt.Println(i, *p)
	fmt.Println(&i, p)
}
