package maths

import (
	"fmt"
)

func main() {

	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
