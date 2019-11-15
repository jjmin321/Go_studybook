//25-4

package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	//defer 함수 (Panic 호출 되면 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(f.Name())
	}
	defer f.Close()
}

func main() {
	//GO panic&recover 최종 실습

	fileOpen("undefined.txt")
	fmt.Println("End Main")

}
