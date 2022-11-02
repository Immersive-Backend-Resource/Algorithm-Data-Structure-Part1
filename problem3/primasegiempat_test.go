package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimaSegiEmpat(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expeted := "17\t19\n23\t29\n31\t37\n156"
		actual := PrimaSegiEmpat(2, 3, 13)
		assert.Equal(t, expeted, actual, "hasil output tidak sesuai")
	})

	t.Run("case 2", func(t *testing.T) {
		expeted := "2\t3\t5\t7\t11\n13\t17\t19\t23\t29\n129"
		actual := PrimaSegiEmpat(5, 2, 1)
		assert.Equal(t, expeted, actual, "hasil output tidak sesuai")
	})
}
