package lc0046_test

import (
	"app/lc0046"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslateNum(t *testing.T) {
	assert.Equal(t, 5, lc0046.TranslateNum(12258))
	assert.Equal(t, 1, lc0046.TranslateNum(0))
	assert.Equal(t, 5, lc0046.TranslateNum(2222))
}
