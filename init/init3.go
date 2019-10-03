package main

import (
	"fmt"
	"package/lib"
)

var num int32

//변수 초기화
func init() {
	num = 30
	fmt.Println("Main init start!")
}

func main() {
	fmt.Println("10보다 큰 수? : ", lib.CheckNum(num))
}
