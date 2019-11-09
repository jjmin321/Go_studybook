//4-2

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//ex1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 or more and less than 100")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, "  25 or more and less than 50")
	default:
		fmt.Println("i -> ", i, " default value")
	}
}
