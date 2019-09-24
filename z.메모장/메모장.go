/*		포인터 기초기초기초기초

package main

import "fmt"

func main() {
	a := 1
	var b *int
	b = &a
	a++
	fmt.Print(*b)	// 2
}

*/

/* 배열복사 배열복사 배열복사 배열복사 배열복사

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	for i := 0; i < 5; i++ {
		b[i] = a[i]
	}
	fmt.Print(b)
}

*/

/*
	 slice는 structure이다.	slice는 3개의 properties(파일확장자)를 가지고 있다.
	slice struct
	pointer *	pointer 라는 시작 주소와
	len int		len 라는 갯수와(사용하고 있는 칸)
	cap int		cap 라는 최대갯수(보유하고 있는 칸)를 가지고 있다.

*/
/*package main

import "fmt"

func main() {
	x := 2
	y := x
	x++
	fmt.Println(y)

	a := 2
	b := &a
	a++
	fmt.Print(*b)
}*/
/*
package main

import "fmt"

func main() {
	var m = map[string]string{
		"jeongmin":  "1asdfasdfasdfasdf",
		"sungwan":   "2asdfasdfasdf",
		"seongheon": "3asdfasdfa",
		"seonghun":  "4asdafasdf",
	}
	fmt.Print(m["jeongmin"])
}*/
