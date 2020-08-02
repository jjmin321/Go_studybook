<div align="center">
  
![GitHub contributors](https://img.shields.io/github/contributors/jjmin321/Go_studybook)
![GitHub forks](https://img.shields.io/github/forks/jjmin321/Go_studybook?label=Forks)
![GitHub stars](https://img.shields.io/github/stars/jjmin321/Go_studybook?style=Stars)
![GitHub issues](https://img.shields.io/github/issues-raw/jjmin321/Go_studybook)
[![Go Report Card](https://goreportcard.com/badge/github.com/jjmin321/Go_studybook)](https://goreportcard.com/report/github.com/jjmin321/Go_studybook)

</div>

# ☁️ Go 언어 특징
1. 컴파일 언어이지만 컴파일러가 소스 코드를 해석하는 pass 수를 줄여서 인터프리터 언어처럼 빠르게 동작합니다.
2. 언어의 문법이 간결하여 접근하기 쉽고 높은 성능을 낼 수 있습니다.
3. 자료형 체계에서 정적 타입 검사가 이루어지기 때문에 Python 등에 익숙해져 있는 경우 생소할 수 있지만 풍부한 라이브러리를 통해서 다양한 기능을 쉽게 구현할 수 있습니다.
4. 고루틴이라는 비동기 매커니즘을 제공하여 이벤트 처리 및 병렬 프로그래밍 작성이 쉽습니다.
5. 고루틴은 OS에서 관리하는 경량 스레드보다 더 가볍기 때문에 CPU 코어갯수와 무관하게 수백, 수천만 고루틴을 작성해도 성능에 문제가 발생하지 않습니다.
6. 파일 언어인 덕분에, 속도가 느린 스크립트 언어에서 연산 퍼포먼스가 필요한 부분을 Go로 대체해 넣을 수도 있습니다.

# 🔑 Go 키워드 총 25가지
break, default, func, interface, select, case, defer, go, map, struct, chan, else, goto, package, switch, const, fallthrough, if, range, type, continue, for, import, return, var

# 📖 Go 터미널 명령어
- go build : 소스 파일 자체의 정보만을 사용하여 Go 바이너리를 빌드한다. 별도의 makefile은 없다.
- go test : 유닛 테스트 및 마이크로벤치마크
- go fmt : 코드 서식 지정
- go get : 원격 패키지의 검색 및 설치
- go vet : 코드 내의 잠재적인 오류를 찾아내는 정적 분석기
- go run : 코드를 빌드하고 실행하는 바로 가기
- godoc : HTTP를 통해 문서 확인
- gorename : 변수, 함수 등을 type-safe 방식으로 이름 변경
- go generate : 코드 생성기를 호출하는 표준 방식

# Go로 사용한 서비스
- Google
- NetFlix
- Twitch
- Twitter
- Dropbox
- Uber
- Docker
- MongoDB
- Paypal

# 🚀 Go_studybook 

🔖 2019년 4월에 (https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65) 게시물을 보고 나서 Go언어 공부를 시작하게 되었습니다.

📁 모든 파일은 공부하면서 다른 입문자들의 GO언어 공부에 도움되었으면 좋겠다는 마음으로 정리해둔 파일입니다.

👨‍💻 GO_Files/StudyBook - GO언어에 존재하는 총 25개의 문법, 파일 입출력에 대한 다양한 예제들을 정리해뒀습니다.

# 👨‍🎓 Go를 사용하는 법 👨‍🎓

1. 기초 : 슬라이스를 비롯한 고 특유의 자료형 접근 방법을 익히고 전통 언어의 객체지향 
방식을 구조체와 메소드 바인딩을 통한 구현 방법을 익힌다.

2. 채널, 동기화, 고루틴 사용  : 채널과 동기화 그리고 고루틴 같은 문법을 이용해 좀 더 Go스러운 프로그램을
구성하고 이런 문법들을 통해 웹 크롤러나 서버를 구축할 때 안정적이고 빠른 프로그램을 만들어본다.

3. 다양한 라이브러리 사용 : yhat/scrape 같은 크롤링 라이브러리부터 웹 서버를 구축하게 도와주는 net/http 라이브러리 등 다양한 라이브러리를 사용해본다.

4. 나만의 모듈 구현 : 이때까지 배운 다양한 문법과 라이브러리를 통해 나만의 모듈을 만들고 사용해본다.

5. 서버 구축 : revel, gin , echo와 같은 서버 프레임워크를 사용하여 백엔드를 구축해본다.

# 2019년 9월 학교에 신청한 도서들 
- GO 마스터하기
- GO언어 실전 테크닉   
- 개발자를 위한 하룻밤에 읽는 GO 이야기
- 디스커버리 GO언어
