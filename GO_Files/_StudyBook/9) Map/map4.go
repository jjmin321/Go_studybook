package main

import "fmt"

func main() {
	//맵(Map)
	//맵 조회할 경우에 주의 할 점

	//ex1
	map1 := map[string]int{ //키가 존재하지 않는다면 int : 0, string : "", float : 0.0으로 예외처리 대신 대입해준다.
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, a := map1["orange"] // 값, 키 존재여부 := 맵이름[키이름]
	value2 := map1["kiwi"]      // 값 := 맵이름[키이름]
	value3, ok := map1["kiwi"]  // 값, 키 존재여부 := 맵이름[키이름]

	fmt.Println(value1, a) //출력(값, 키 존재여부)
	fmt.Println(value2)
	fmt.Println(value3, ok) //두번째 리턴 값으로 키 존재 유무 확인
	fmt.Println()
	//ex2
	i := "kiwi is not exist!!!"

	if _, ok := map1["kiwi"]; ok {
	} else {
		fmt.Println(i)
	}
	// same same same same same same same same same
	if _, ok := map1["kiwi"]; !ok { // 만약 map1["kiwi"]이라는 키가 존재하지 않는다면 kiwi is not exist!!!를 출력한다.
		fmt.Println(i)
	}

}
