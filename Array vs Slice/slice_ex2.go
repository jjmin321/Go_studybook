package main

import (
	"fmt"
	"sort"
)

func main() {
	//슬라이스 추출 및 정렬
	//slice[i:j] i부터 j-1까지 추출
	//slice[i:] i부터 끝까지 추출
	//slice[:j] 처음부터 j까지 추출
	//slice[:] 처음부터 끝까지 추출

	//ex1(추출)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex1 : ", slice1[:])
	fmt.Println("ex1 : ", slice1[0:])
	fmt.Println("ex1 : ", slice1[:5])
	fmt.Println("ex1 : ", slice1[0:len(slice1)])
	fmt.Println("ex1 : ", slice1[3:])
	fmt.Println("ex1 : ", slice1[1:3])

	//ex2(sort)
	//sort package : https://golang.org/pkg/sort

	slice2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	slice3 := []string{`B`, `D`, `F`, `A`, `C`, `E`}

	//sort하는 법 !!!!!!!!!!!!!! important important important important important
	fmt.Println()
	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2)) //sort 돼있는지 확인하는법
	sort.Ints(slice2)                                 //sort 하는 법
	fmt.Println("ex2 : ", slice2)
	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2))
	fmt.Println()
	fmt.Println("ex3 : ", sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println("ex3 : ", slice3)
	fmt.Println("ex3 : ", sort.StringsAreSorted(slice3))

}
