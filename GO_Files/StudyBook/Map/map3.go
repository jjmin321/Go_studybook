//9-3

package main

import "fmt"

func main() {
	//맵(Map)
	//맵 값 변경 및 삭제

	//ex1
	map1 := map[string]string{ // map1 := make(map[string]string)
		"daum":   "http://daum.net", // map1["daum"] = "http://daum.net"
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://home1.com",
	}
	fmt.Println(map1)

	map1["home2"] = "http://home2.com" //add
	fmt.Println(map1)

	map1["home2"] = "http://test2.com" //modify
	fmt.Println(map1)

	//ex2 (delete)
	delete(map1, "home2") //delete
	fmt.Println(map1)

	delete(map1, "daum") //delete
	fmt.Println(map1)

	delete(map1, "naver") //delete
	fmt.Println(map1)

	delete(map1, "google") //delete
	fmt.Println(map1)

	delete(map1, "home1") //delete
	fmt.Println(map1)

}
