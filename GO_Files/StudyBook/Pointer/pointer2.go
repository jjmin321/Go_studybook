//10-2

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
