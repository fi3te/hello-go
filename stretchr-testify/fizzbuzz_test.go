package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	a := assert.New(t)

	_, err := fizzBuzz(0)
	a.NotNil(err)

	value, err := fizzBuzz(1)
	a.Nil(err)
	a.Equal("1", value)

	value, err = fizzBuzz(3)
	a.Nil(err)
	a.Equal("Fizz", value)

	value, err = fizzBuzz(5)
	a.Nil(err)
	a.Equal("Buzz", value)

	value, err = fizzBuzz(15)
	a.Nil(err)
	a.Equal("Fizzbuzz", value)
}
