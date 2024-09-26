package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(value int) (string, error) {
	if value < 1 {
		return "", fmt.Errorf("%v is not a valid input for fizzBuzz", value)
	}
	fizz := value%3 == 0
	buzz := value%5 == 0
	if fizz && buzz {
		return "Fizzbuzz", nil
	} else if fizz {
		return "Fizz", nil
	} else if buzz {
		return "Buzz", nil
	} else {
		return strconv.Itoa(value), nil
	}
}
