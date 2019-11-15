//26-1

package main

import (
	"fmt"
	"os"
)

//에러 체크 방식1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

//에러 체크 방식2
func errCheck2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//파일 쓰기
	//Create : 새 파일 작성 및 파일 열기 기능
	//Close : 리소스 닫기
	//Write, WriteString, WriteAt : 파일 쓰기
	//각 운영체제 권한 주의(오류 메시지 확인)
	//예외 처리 정말 중요!
	//https://golang.org/pkg/os

	//파일 쓰기 ex1
	file, err := os.Create("Users/jejeongmin/documents/go/src/go_studybook/go_files/studybook/filewrite/test_write.txt") //이 경로로 파일 생성
	errCheck1(err)

	defer file.Close()

	//쓰기 ex1
	s1 := []byte{77, 78, 79, 80, 81}
	n1, err := file.Write([]byte(s1)) //문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	errCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n", n1)

	file.Sync() //Write Commit(Stable)!

	//쓰기 ex2
	s2 := "Hello Golang!\n File Write Test! - 1\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("쓰기 작업(2) 완료 (%d bytes) \n", n2)

	file.Sync() //없어도 되지만 있으면 깔끔하고 간지 !

	//쓰기 ex3
	s3 := "Test WriteAt! -2\n"
	n3, err := file.WriteAt([]byte(s3), 50) //len(offset) 조절하면서 테스트
	errCheck1(err)
	fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n", n3)

	file.Sync()

	//쓰기 ex4
	n4, err := file.WriteString("\nHello Golang! \n File Write Test! - 3\n")
	errCheck1(err)
	fmt.Printf("쓰기 작업(4) 완료 (%d bytes) \n", n4)

}
