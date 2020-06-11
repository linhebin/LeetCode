package lc0009_test

import (
	"app/lc0009"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, false, lc0009.IsPalindrome(12))
	assert.Equal(t, true, lc0009.IsPalindrome(121))
}

func TestIsPalindrome2(t *testing.T) {
	assert.Equal(t, false, lc0009.IsPalindrome2(12))
	assert.Equal(t, true, lc0009.IsPalindrome2(61216))
}
