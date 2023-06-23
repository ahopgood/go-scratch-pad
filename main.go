package main

import (
	"com/alexander/scratch/maths"
	"fmt"
)

func main() {

	f := maths.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
