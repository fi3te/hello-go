package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		value, err := fizzBuzz(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}
}
