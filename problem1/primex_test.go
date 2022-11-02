package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeX(t *testing.T) {
	t.Run("Case 1 : parameter bernilai 1", func(t *testing.T) {
		expected := 2
		actual := PrimeX(1)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("Case 2 : parameter bernilai 5", func(t *testing.T) {
		expected := 11
		actual := PrimeX(5)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("Case 3 : parameter bernilai 8", func(t *testing.T) {
		expected := 19
		actual := PrimeX(8)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("Case 4 : parameter bernilai 9", func(t *testing.T) {
		expected := 23
		actual := PrimeX(9)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

	t.Run("Case 5 : parameter bernilai 10", func(t *testing.T) {
		expected := 29
		actual := PrimeX(10)
		assert.Equal(t, expected, actual, "hasil output tidak sesuai")
	})

}
