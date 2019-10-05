package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//다른 타입끼리는 반드시 형 변환 후 연산
	//형 변환이 없을 경우 예외(에러) 발생
	// +, -, *, %, /, <<, >>, &, ^

	//ex1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1 : ", n1)
	fmt.Println("ex2 : ", n2)
	fmt.Println("ex3 : ", n3)
	fmt.Println("ex4 : ", n4)
	fmt.Println("ex5 : ", math.MaxFloat32)
	fmt.Println("ex6 : ", math.MaxInt16)
	fmt.Println("ex7 : ", math.MaxInt32)

	n5 := 100000 //int
	n6 := int16(10000)

	fmt.Println("ex8 : ", int16(n5))
	fmt.Println("ex9 : ", int16(n5) > n6)
	fmt.Println("ex10 : ", n6-int16(n5))
}
