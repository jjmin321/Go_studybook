package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	//포인터 값 전달
	//함수, 메서드 호출 시에 매개변수 값을 복수 전달 -> 함수, 메서드 내에서는 원본 값 변경 불가능
	//원본 값 변경 위해서 포인터 전달
	//특히 크기가 큰 배열인 경우 값 복사시에 시스템 부담 -> 포인터 전달로 해결(슬라이스, 맵 참조 전달) [배열에 포인터를 붙히면 슬라이스처럼 사용 가능]

	//ex1
	var a int = 10
	var b int = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()
	rptc(&a) //포인터형에는 주소값이 들어가야 하므로 &a로 전달
	vptc(b)
	fmt.Println(a)
	fmt.Println(b)
}
