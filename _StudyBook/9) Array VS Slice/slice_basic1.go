package main

import "fmt"

func main() {
	//길이가 고정되어 있지 않고 늘어난다 (가변) -> reference type
	//슬라이스 (길이 & 용량) 크기가 동적으로 할당 가능
	//배열 vs 슬라이스 차이점 중요 important important important
	//길이 고정(상수와 비슷) vs 길이 가변(변수와 비슷)
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	//비교 연산자 사용 가능 vs 비교 연산자 사용 불가 (주소값이라서)
	//대부분 슬라이스 사용한다.
	//cap() : 배열, 슬라이스 용량
	//len() : 배열 , 슬라이스 길이
	//2가지 선언 방법 1. 배열처럼 선언, 2. make함수 : make(type, len, cap(생략 시 길이))

	//ex1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	//slice2[2] = 10 //slice는 가변형이라 값이 초기화되지 않았으면 배열처럼 값을 못 바꿈.
	slice3[4] = 5 //값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	//ex2
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100) //best thing
	slice8 := make([]int, 5)

	slice6[2] = 7 //삽입
	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	//ex3
	var slice9 []int //nil 슬라이스 (길이와 용량 0)

	if slice9 == nil {
		fmt.Println("This is nil slice!")
	}
}
