package lc0014_test

import (
	"app/lc0014"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"dog2","dog","dog1"}
	comStr := lc0014.LongestCommonPrefix(strs)
	assert.Equal(t, "dog", comStr)
}