//28-1

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//파일 읽기, 쓰기 -> ioutil 패키지 활용
	//더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	//아래 메소드 확인 가능
	//WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	//https://golang.org/pkg/io/ioutil

	s := "Hello Golang!\n File Write Test!\n"

	//파일모드 -> 퍼미션
	//읽기 : 4, 쓰기 : 2, 실행 : 1
	//소유자, 그룹, 기타 사용자 순서 (777 -> 모두 읽기, 쓰기, 실행 가능)
	//https://golang.org/pkg/os/#FileMode
	err := ioutil.WriteFile("Users/jejeongmin/documents/go/src/go_studybook/go_files/studybook/fileioutil/test_write1.txt", []byte(s), os.FileMode(0444))
	// go에서는 퍼미션(소유자 , 그룹, 기타 사용자 순서로 읽기, 쓰기, 실행 권한)을 줄 때 앞에 0붙힘.
	errCheck(err)

	data, err := ioutil.ReadFile("users/jejeongmin/documents/go/src/go_studybook/go_files/studybook/fileioutil/test_write1.txt")
	errCheck(err)
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")
	fmt.Println("test_write1.txt : ", string(data))
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")

}
