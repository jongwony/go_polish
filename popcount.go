package popcount

import (
	"fmt"
)

var a = b + c
var b = f()
var c = 1

func f() int { return c + 1 }

func main() {
	fmt.Println(f())
}
