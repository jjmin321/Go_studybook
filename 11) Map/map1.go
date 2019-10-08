package main

import "fmt"

func main() {
	//맵(Map)
	//맵: 해시테이블, 딕셔너리(파이썬) , Key - Value로 자료 저장
	//레퍼런스 타입(참조 값 전달)
	//참조 타입이므로 비교 연산자 사요 불가능
	//특징 : 참조타입(Key)로 사용이 불가능, 값으로는 모든 타입 사용 가능
	//make 함수 및 축약(리터럴)로 초기화 가능
	//순서 없음

	//ex1
	var map1 map[string]int = make(map[string]int) //정석
	var map2 = make(map[string]int)                //자료형 생략
	map3 := make(map[string]int)                   //리터럴 형 (comfortable)

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println()

	//ex2
	// map4 := map[string]int{} //json 형태
	map4 := make(map[string]int)
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"orange": 23,
		"apple":  15,
		"banana": 40,
	}

	map6 := make(map[string]int, 10)
	map6["orange"] = 33
	map6["apple"] = 22
	map6["banana"] = 44

	fmt.Println(map4)
	fmt.Println(map5)
	fmt.Println(map6)
	fmt.Println(map6["banana"])
	fmt.Println(map6["orange"])
}
