package lc0125_test

import (
	"app/lc0125"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	//assert.Equal(t, true, lc0125.IsPalindrome("A man, a plan, a canal: Panama"))
	//assert.Equal(t, false, lc0125.IsPalindrome("race a car"))
	assert.Equal(t, false, lc0125.IsPalindrome("AZaz09"))
}
