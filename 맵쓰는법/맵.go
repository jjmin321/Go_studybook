/*package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Hello world!"] = 4231
	fmt.Print(m["Hello world!"])
}
*/

package main

import "fmt"

var m = map[string]int{
	"jeongmin":  12345,
	"sungwan":   678910,
	"seongheon": 1345,
	"seonghun":  235,
}

func main() {
	fmt.Println(m["sungwan"])
}
