package lc0739_test

import (
	"app/lc0739"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	a := []int{1, 1, 4, 2, 1, 1, 0, 0}
	b := []int{73, 74, 75, 71, 69, 72, 76, 73}
	c := lc0739.DailyTemperatures(b)
	assert.Equal(t, a, c)
}
