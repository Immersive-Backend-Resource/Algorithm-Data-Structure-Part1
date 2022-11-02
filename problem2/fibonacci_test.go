package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {

	t.Run("case 1 : parameter bernilai 0", func(t *testing.T) {
		expected := 0
		actual := Fibonacci(0)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("case 2 : parameter bernilai 2", func(t *testing.T) {
		expected := 1
		actual := Fibonacci(2)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("case 3 : parameter bernilai 9", func(t *testing.T) {
		expected := 34
		actual := Fibonacci(9)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("case 4 : parameter bernilai 10", func(t *testing.T) {
		expected := 55
		actual := Fibonacci(10)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("case 5 : parameter bernilai 12", func(t *testing.T) {
		expected := 144
		actual := Fibonacci(12)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

}
