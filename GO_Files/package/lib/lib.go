//라이브러리 접근제어(1)
package lib //go/src/어플리케이션/패키지폴더 경로로 만든 후 import "어플리케이션/패키지폴더" 로 패키지파일 사용 가능

import "fmt"

func init() {
	fmt.Println("lib init start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
