//8-8

package main

import "fmt"

func main() {
	//slice copy
	//copy (복사 대상, 원본)
	//반드시 make로 공간을 할당 후 복사해야 복사가 된다.
	//복사 된 슬라이스 값 변경해도 원본에는 영향 없음 [대참사를 방지하기 위해 복사를 해서 테스트해보고 원본에 옮길 때 씀]

	//ex1 copy
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1) //공간이 없어서 복사가 안 됨.

	fmt.Println(slice2)
	fmt.Println(slice3)

	//ex2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)

	b[0] = 7
	b[4] = 10
	fmt.Println(a) //b는 a의 복사품이므로 a는 값이 바뀌지 않음
	fmt.Println(b)
	fmt.Println()

	//ex3
	c := make([]int, 5, 10)
	c = []int{1, 2, 3, 4, 5}
	d := make([]int, 5, 5)

	copy(d, c[0:3])

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println()

	//ex4
	e := make([]int, 10, 10)
	e = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]
	fmt.Println(e)
	fmt.Println(f)

}
