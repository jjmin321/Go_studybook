package main

import "fmt"

func main() {
	//배열
	//배열은 용량 , 길이 항상 같다.
	//배열 vs 슬라이스 차이점 중요 important important important
	//길이 고정(상수와 비슷) vs 길이 가변(변수와 비슷)
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	//비교 연산자 사용 가능 vs 비교 연산자 사용 불가 (주소값이라서)
	//대부분 슬라이스 사용한다.

	//cap() : 배열, 슬라이스 용량
	//len() : 배열 , 슬라이스 길이

	//ex1
	var arr1 [5]int
	var arr2 [5]int = [5]int{0, 1, 2, 3, 4}
	var arr3 = [5]int{0, 1, 2, 3, 4}
	arr4 := [5]int{0, 1, 2, 3, 4}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3}
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	arr1[2] = 5

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr7), arr7)

	//ex2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}

	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

}
